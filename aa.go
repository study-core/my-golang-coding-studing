package main

import (
	"RosettaFlow/carrier-go/common/timeutils"
	"strings"

	//"RosettaFlow/carrier-go/types"
	"context"
	"math"
	"sync"

	//"container/heap"
	//"context"
	"fmt"
	"time"

	//"math"
	//"time"
)

//func main() {
//
//	var x int
//	threads := runtime.GOMAXPROCS(0)
//	for i := 0; i < threads; i++ {
//		go func() {
//
//			for {
//				x++
//			}
//
//		}()
//
//	}
//	time.Sleep(time.Second)
//	fmt.Println("x =", x)
//
//}

func main() {

	//var x int
	//
	//threads := runtime.GOMAXPROCS(0) - 1
	//
	//for i := 0; i < threads; i++ {
	//	go func() {
	//		for {
	//			x++
	//		}
	//	}()
	//}
	//time.Sleep(time.Second)
	//fmt.Println("x =", x)

	//fmt.Println(fA())
	//fmt.Println(fB())
	//fmt.Println(fC())

	//flag := true
	//
	//fmt.Printf("flag: %v\n", flag)

	// 训练 py
	//train_code := "# coding:utf-8\n\nimport sys\nsys.path.append(\"..\")\nimport os\nimport math\nimport json\nimport time\nimport logging\nimport shutil\nimport numpy as np\nimport pandas as pd\nimport tensorflow as tf\nimport latticex.rosetta as rtt\nimport channel_sdk\n\n\nnp.set_printoptions(suppress=True)\ntf.compat.v1.logging.set_verbosity(tf.compat.v1.logging.ERROR)\nos.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'\nrtt.set_backend_loglevel(5)  # All(0), Trace(1), Debug(2), Info(3), Warn(4), Error(5), Fatal(6)\nlog = logging.getLogger(__name__)\n\nclass PrivacyLRTrain(object):\n    '''\n    Privacy logistic regression train base on rosetta.\n    '''\n\n    def __init__(self,\n                 channel_config: str,\n                 cfg_dict: dict,\n                 data_party: list,\n                 result_party: list,\n                 results_dir: str):\n        log.info(f\"channel_config:{channel_config}, cfg_dict:{cfg_dict}, data_party:{data_party}, \"\n                 f\"result_party:{result_party}, results_dir:{results_dir}\")\n        assert isinstance(channel_config, str), \"type of channel_config must be str\"\n        assert isinstance(cfg_dict, dict), \"type of cfg_dict must be dict\"\n        assert isinstance(data_party, (list, tuple)), \"type of data_party must be list or tuple\"\n        assert isinstance(result_party, (list, tuple)), \"type of result_party must be list or tuple\"\n        assert isinstance(results_dir, str), \"type of results_dir must be str\"\n        \n        self.channel_config = channel_config\n        self.data_party = list(data_party)\n        self.result_party = list(result_party)\n        self.party_id = cfg_dict[\"party_id\"]\n        self.input_file = cfg_dict[\"data_party\"].get(\"input_file\")\n        self.key_column = cfg_dict[\"data_party\"].get(\"key_column\")\n        self.selected_columns = cfg_dict[\"data_party\"].get(\"selected_columns\")\n\n        dynamic_parameter = cfg_dict[\"dynamic_parameter\"]\n        self.label_owner = dynamic_parameter.get(\"label_owner\")\n        if self.party_id == self.label_owner:\n            self.label_column = dynamic_parameter.get(\"label_column\")\n            self.data_with_label = True\n        else:\n            self.label_column = \"\"\n            self.data_with_label = False\n                        \n        algorithm_parameter = dynamic_parameter[\"algorithm_parameter\"]\n        self.epochs = algorithm_parameter.get(\"epochs\", 10)\n        self.batch_size = algorithm_parameter.get(\"batch_size\", 256)\n        self.learning_rate = algorithm_parameter.get(\"learning_rate\", 0.001)\n        self.use_validation_set = algorithm_parameter.get(\"use_validation_set\", True)\n        self.validation_set_rate = algorithm_parameter.get(\"validation_set_rate\", 0.2)\n        self.predict_threshold = algorithm_parameter.get(\"predict_threshold\", 0.5)\n\n        self.output_file = os.path.join(results_dir, \"model\")\n        \n        self.check_parameters()\n\n    def check_parameters(self):        \n        assert self.epochs > 0, \"epochs must be greater 0\"\n        assert self.batch_size > 0, \"batch size must be greater 0\"\n        assert self.learning_rate > 0, \"learning rate must be greater 0\"\n        assert 0 < self.validation_set_rate < 1, \"validattion set rate must be between (0,1)\"\n        assert 0 <= self.predict_threshold <= 1, \"predict threshold must be between [0,1]\"\n        \n    def train(self):\n        '''\n        Logistic regression training algorithm implementation function\n        '''\n\n        log.info(\"extract feature or label.\")\n        train_x, train_y, val_x, val_y = self.extract_feature_or_label(with_label=self.data_with_label)\n        \n        log.info(\"start create and set channel.\")\n        self.create_set_channel()\n        log.info(\"waiting other party connect...\")\n        rtt.activate(\"SecureNN\")\n        log.info(\"protocol has been activated.\")\n        \n        log.info(f\"start set save model. save to party: {self.result_party}\")\n        rtt.set_saver_model(False, plain_model=self.result_party)\n        # sharing data\n        log.info(f\"start sharing train data. data_owner={self.data_party}, label_owner={self.label_owner}\")\n        shard_x, shard_y = rtt.PrivateDataset(data_owner=self.data_party, label_owner=self.label_owner).load_data(train_x, train_y, header=0)\n        log.info(\"finish sharing train data.\")\n        column_total_num = shard_x.shape[1]\n        \n        if self.use_validation_set:\n            log.info(\"start sharing validation data.\")\n            shard_x_val, shard_y_val = rtt.PrivateDataset(data_owner=self.data_party, label_owner=self.label_owner).load_data(val_x, val_y, header=0)\n            log.info(\"finish sharing validation data.\")\n\n        if self.party_id not in self.data_party:  \n            # mean the compute party and result party\n            log.info(\"compute start.\")\n            X = tf.placeholder(tf.float64, [None, column_total_num])\n            Y = tf.placeholder(tf.float64, [None, 1])\n            W = tf.Variable(tf.zeros([column_total_num, 1], dtype=tf.float64))\n            b = tf.Variable(tf.zeros([1], dtype=tf.float64))\n            logits = tf.matmul(X, W) + b\n            loss = tf.nn.sigmoid_cross_entropy_with_logits(labels=Y, logits=logits)\n            loss = tf.reduce_mean(loss)\n            # optimizer\n            optimizer = tf.train.GradientDescentOptimizer(self.learning_rate).minimize(loss)\n            init = tf.global_variables_initializer()\n            saver = tf.train.Saver(var_list=None, max_to_keep=5, name='v2')\n            \n            pred_Y = tf.sigmoid(tf.matmul(X, W) + b)\n            reveal_Y = rtt.SecureReveal(pred_Y)\n            actual_Y = tf.placeholder(tf.float64, [None, 1])\n            reveal_Y_actual = rtt.SecureReveal(actual_Y)\n\n            with tf.Session() as sess:\n                log.info(\"session init.\")\n                sess.run(init)\n                # train\n                log.info(\"train start.\")\n                train_start_time = time.time()\n                batch_num = math.ceil(len(shard_x) / self.batch_size)\n                for e in range(self.epochs):\n                    for i in range(batch_num):\n                        bX = shard_x[(i * self.batch_size): (i + 1) * self.batch_size]\n                        bY = shard_y[(i * self.batch_size): (i + 1) * self.batch_size]\n                        sess.run(optimizer, feed_dict={X: bX, Y: bY})\n                        if (i % 50 == 0) or (i == batch_num - 1):\n                            log.info(f\"epoch:{e + 1}/{self.epochs}, batch:{i + 1}/{batch_num}\")\n                log.info(f\"model save to: {self.output_file}\")\n                saver.save(sess, self.output_file)\n                train_use_time = round(time.time()-train_start_time, 3)\n                log.info(f\"save model success. train_use_time={train_use_time}s\")\n                \n                if self.use_validation_set:\n                    Y_pred = sess.run(reveal_Y, feed_dict={X: shard_x_val})\n                    log.debug(f\"Y_pred:\\n {Y_pred[:10]}\")\n                    Y_actual = sess.run(reveal_Y_actual, feed_dict={actual_Y: shard_y_val})\n                    log.debug(f\"Y_actual:\\n {Y_actual[:10]}\")\n        \n            running_stats = str(rtt.get_perf_stats(True)).replace('\\n', '').replace(' ', '')\n            log.info(f\"running stats: {running_stats}\")\n        else:\n            log.info(\"computing, please waiting for compute finish...\")\n        rtt.deactivate()\n     \n        log.info(\"remove temp dir.\")\n        if self.party_id in (self.data_party + self.result_party):\n            self.remove_temp_dir()\n        else:\n            # delete the model in the compute party.\n            self.remove_output_dir()\n        \n        if (self.party_id in self.result_party) and self.use_validation_set:\n            log.info(\"result_party evaluate model.\")\n            from sklearn.metrics import roc_auc_score, roc_curve, f1_score, precision_score, recall_score, accuracy_score\n            Y_pred_prob = Y_pred.astype(\"float\").reshape([-1, ])\n            Y_true = Y_actual.astype(\"float\").reshape([-1, ])\n            auc_score = roc_auc_score(Y_true, Y_pred_prob)\n            Y_pred_class = (Y_pred_prob > self.predict_threshold).astype('int64')  # default threshold=0.5\n            accuracy = accuracy_score(Y_true, Y_pred_class)\n            f1_score = f1_score(Y_true, Y_pred_class)\n            precision = precision_score(Y_true, Y_pred_class)\n            recall = recall_score(Y_true, Y_pred_class)\n            log.info(\"********************\")\n            log.info(f\"AUC: {round(auc_score, 6)}\")\n            log.info(f\"ACCURACY: {round(accuracy, 6)}\")\n            log.info(f\"F1_SCORE: {round(f1_score, 6)}\")\n            log.info(f\"PRECISION: {round(precision, 6)}\")\n            log.info(f\"RECALL: {round(recall, 6)}\")\n            log.info(\"********************\")\n        log.info(\"train finish.\")\n    \n    def create_set_channel(self):\n        '''\n        create and set channel.\n        '''\n        io_channel = channel_sdk.grpc.APIManager()\n        log.info(\"start create channel\")\n        channel = io_channel.create_channel(self.party_id, self.channel_config)\n        log.info(\"start set channel\")\n        rtt.set_channel(\"\", channel)\n        log.info(\"set channel success.\")\n    \n    def extract_feature_or_label(self, with_label: bool=False):\n        '''\n        Extract feature columns or label column from input file,\n        and then divide them into train set and validation set.\n        '''\n        train_x = \"\"\n        train_y = \"\"\n        val_x = \"\"\n        val_y = \"\"\n        temp_dir = self.get_temp_dir()\n        if self.party_id in self.data_party:\n            if self.input_file:\n                if with_label:\n                    usecols = self.selected_columns + [self.label_column]\n                else:\n                    usecols = self.selected_columns\n                \n                input_data = pd.read_csv(self.input_file, usecols=usecols, dtype=\"str\")\n                input_data = input_data[usecols]\n                # only if self.validation_set_rate==0, split_point==input_data.shape[0]\n                split_point = int(input_data.shape[0] * (1 - self.validation_set_rate))\n                assert split_point > 0, f\"train set is empty, because validation_set_rate:{self.validation_set_rate} is too big\"\n                \n                if with_label:\n                    y_data = input_data[self.label_column]\n                    train_y = os.path.join(temp_dir, f\"train_y_{self.party_id}.csv\")\n                    y_data.iloc[:split_point].to_csv(train_y, header=True, index=False)\n                    if self.use_validation_set:\n                        assert split_point < input_data.shape[0], \\\n                            f\"validation set is empty, because validation_set_rate:{self.validation_set_rate} is too small\"\n                        val_y = os.path.join(temp_dir, f\"val_y_{self.party_id}.csv\")\n                        y_data.iloc[split_point:].to_csv(val_y, header=True, index=False)\n                    del input_data[self.label_column]\n                \n                x_data = input_data\n                train_x = os.path.join(temp_dir, f\"train_x_{self.party_id}.csv\")\n                x_data.iloc[:split_point].to_csv(train_x, header=True, index=False)\n                if self.use_validation_set:\n                    assert split_point < input_data.shape[0], \\\n                            f\"validation set is empty, because validation_set_rate:{self.validation_set_rate} is too small\"\n                    val_x = os.path.join(temp_dir, f\"val_x_{self.party_id}.csv\")\n                    x_data.iloc[split_point:].to_csv(val_x, header=True, index=False)\n            else:\n                raise Exception(f\"data_node {self.party_id} not have data. input_file:{self.input_file}\")\n        return train_x, train_y, val_x, val_y\n    \n    def get_temp_dir(self):\n        '''\n        Get the directory for temporarily saving files\n        '''\n        temp_dir = os.path.join(os.path.dirname(self.output_file), 'temp')\n        if not os.path.exists(temp_dir):\n            os.makedirs(temp_dir, exist_ok=True)\n        return temp_dir\n\n    def remove_temp_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        Only delete temp file.\n        '''\n        temp_dir = self.get_temp_dir()\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n    \n    def remove_output_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        This is used to delete all output files of the non-resulting party\n        '''\n        temp_dir = os.path.dirname(self.output_file)\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n\n\ndef main(channel_config: str, cfg_dict: dict, data_party: list, result_party: list, results_dir: str):\n    '''\n    This is the entrance to this module\n    '''\n    privacy_lr = PrivacyLRTrain(channel_config, cfg_dict, data_party, result_party, results_dir)\n    privacy_lr.train()\n"
	//
	//predict_code := "# coding:utf-8\n\nimport sys\nsys.path.append(\"..\")\nimport os\nimport math\nimport json\nimport time\nimport logging\nimport shutil\nimport numpy as np\nimport pandas as pd\nimport tensorflow as tf\nimport latticex.rosetta as rtt\nimport channel_sdk\n\n\nnp.set_printoptions(suppress=True)\ntf.compat.v1.logging.set_verbosity(tf.compat.v1.logging.ERROR)\nos.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'\nrtt.set_backend_loglevel(5)  # All(0), Trace(1), Debug(2), Info(3), Warn(4), Error(5), Fatal(6)\nlog = logging.getLogger(__name__)\n\nclass PrivacyLRPredict(object):\n    '''\n    Privacy logistic regression predict base on rosetta.\n    '''\n\n    def __init__(self,\n                 channel_config: str,\n                 cfg_dict: dict,\n                 data_party: list,\n                 result_party: list,\n                 results_dir: str):\n        log.info(f\"channel_config:{channel_config}, cfg_dict:{cfg_dict}, data_party:{data_party},\"\n                 f\"result_party:{result_party}, results_dir:{results_dir}\")\n        assert isinstance(channel_config, str), \"type of channel_config must be str\"\n        assert isinstance(cfg_dict, dict), \"type of cfg_dict must be dict\"\n        assert isinstance(data_party, (list, tuple)), \"type of data_party must be list or tuple\"\n        assert isinstance(result_party, (list, tuple)), \"type of result_party must be list or tuple\"\n        assert isinstance(results_dir, str), \"type of results_dir must be str\"\n        \n        self.channel_config = channel_config\n        self.data_party = list(data_party)\n        self.result_party = list(result_party)\n        self.party_id = cfg_dict[\"party_id\"]\n        self.input_file = cfg_dict[\"data_party\"].get(\"input_file\")\n        self.key_column = cfg_dict[\"data_party\"].get(\"key_column\")\n        self.selected_columns = cfg_dict[\"data_party\"].get(\"selected_columns\")\n        dynamic_parameter = cfg_dict[\"dynamic_parameter\"]\n        self.model_restore_party = dynamic_parameter.get(\"model_restore_party\")\n        model_path = dynamic_parameter.get(\"model_path\")\n        self.model_file = f\"{model_path}/model\"\n        self.predict_threshold = dynamic_parameter.get(\"predict_threshold\", 0.5)\n        assert 0 <= self.predict_threshold <= 1, \"predict threshold must be between [0,1]\"\n        \n        self.output_file = os.path.join(results_dir, \"result\")\n        \n        self.data_party.remove(self.model_restore_party)  # except restore party\n       \n\n    def predict(self):\n        '''\n        Logistic regression predict algorithm implementation function\n        '''\n\n        log.info(\"extract feature or id.\")\n        file_x, id_col = self.extract_feature_or_index()\n        \n        log.info(\"start create and set channel.\")\n        self.create_set_channel()\n        log.info(\"waiting other party connect...\")\n        rtt.activate(\"SecureNN\")\n        log.info(\"protocol has been activated.\")\n        \n        log.info(f\"start set restore model. restore party={self.model_restore_party}\")\n        rtt.set_restore_model(False, plain_model=self.model_restore_party)\n        # sharing data\n        log.info(f\"start sharing data. data_owner={self.data_party}\")\n        shard_x = rtt.PrivateDataset(data_owner=self.data_party).load_X(file_x, header=0)\n        log.info(\"finish sharing data .\")\n        column_total_num = shard_x.shape[1]\n\n        X = tf.placeholder(tf.float64, [None, column_total_num])\n        W = tf.Variable(tf.zeros([column_total_num, 1], dtype=tf.float64))\n        b = tf.Variable(tf.zeros([1], dtype=tf.float64))\n        saver = tf.train.Saver(var_list=None, max_to_keep=5, name='v2')\n        init = tf.global_variables_initializer()\n        # predict\n        pred_Y = tf.sigmoid(tf.matmul(X, W) + b)\n        reveal_Y = rtt.SecureReveal(pred_Y)  # only reveal to result party\n\n        with tf.Session() as sess:\n            log.info(\"session init.\")\n            sess.run(init)\n            log.info(\"start restore model.\")\n            if self.party_id == self.model_restore_party:\n                if os.path.exists(os.path.join(os.path.dirname(self.model_file), \"checkpoint\")):\n                    log.info(f\"model restore from: {self.model_file}.\")\n                    saver.restore(sess, self.model_file)\n                else:\n                    raise Exception(\"model not found or model damaged\")\n            else:\n                log.info(\"restore model...\")\n                temp_file = os.path.join(self.get_temp_dir(), 'ckpt_temp_file')\n                with open(temp_file, \"w\") as f:\n                    pass\n                saver.restore(sess, temp_file)\n            log.info(\"finish restore model.\")\n            \n            # predict\n            log.info(\"predict start.\")\n            predict_start_time = time.time()\n            Y_pred_prob = sess.run(reveal_Y, feed_dict={X: shard_x})\n            log.debug(f\"Y_pred_prob:\\n {Y_pred_prob[:10]}\")\n            predict_use_time = round(time.time() - predict_start_time, 3)\n            log.info(f\"predict success. predict_use_time={predict_use_time}s\")\n        rtt.deactivate()\n        log.info(\"rtt deactivate finish.\")\n        \n        if self.party_id in self.result_party:\n            log.info(\"predict result write to file.\")\n            output_file_predict_prob = os.path.splitext(self.output_file)[0] + \"_predict.csv\"\n            Y_pred_prob = Y_pred_prob.astype(\"float\")\n            Y_prob = pd.DataFrame(Y_pred_prob, columns=[\"Y_prob\"])\n            Y_class = (Y_pred_prob > self.predict_threshold) * 1\n            Y_class = pd.DataFrame(Y_class, columns=[\"Y_class\"])\n            Y_result = pd.concat([Y_prob, Y_class], axis=1)\n            Y_result.to_csv(output_file_predict_prob, header=True, index=False)\n        log.info(\"start remove temp dir.\")\n        self.remove_temp_dir()\n        log.info(\"predict finish.\")\n\n    def create_set_channel(self):\n        '''\n        create and set channel.\n        '''\n        io_channel = channel_sdk.grpc.APIManager()\n        log.info(\"start create channel\")\n        channel = io_channel.create_channel(self.party_id, self.channel_config)\n        log.info(\"start set channel\")\n        rtt.set_channel(\"\", channel)\n        log.info(\"set channel success.\")\n        \n    def extract_feature_or_index(self):\n        '''\n        Extract feature columns or index column from input file.\n        '''\n        file_x = \"\"\n        id_col = None\n        temp_dir = self.get_temp_dir()\n        if self.party_id in self.data_party:\n            if self.input_file:\n                usecols = [self.key_column] + self.selected_columns\n                input_data = pd.read_csv(self.input_file, usecols=usecols, dtype=\"str\")\n                input_data = input_data[usecols]\n                id_col = input_data[self.key_column]\n                file_x = os.path.join(temp_dir, f\"file_x_{self.party_id}.csv\")\n                x_data = input_data.drop(labels=self.key_column, axis=1)\n                x_data.to_csv(file_x, header=True, index=False)\n            else:\n                raise Exception(f\"data_party:{self.party_id} not have data. input_file:{self.input_file}\")\n        return file_x, id_col\n    \n    def get_temp_dir(self):\n        '''\n        Get the directory for temporarily saving files\n        '''\n        temp_dir = os.path.join(os.path.dirname(self.output_file), 'temp')\n        if not os.path.exists(temp_dir):\n            os.makedirs(temp_dir, exist_ok=True)\n        return temp_dir\n\n    def remove_temp_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        Only delete temp file.\n        '''\n        temp_dir = self.get_temp_dir()\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n\n\ndef main(channel_config: str, cfg_dict: dict, data_party: list, result_party: list, results_dir: str):\n    '''\n    This is the entrance to this module\n    '''\n    privacy_lr = PrivacyLRPredict(channel_config, cfg_dict, data_party, result_party, results_dir)\n    privacy_lr.predict()\n"

	//
	//train_code := "# coding:utf-8\n\nimport sys\nsys.path.append(\"..\")\nimport os\nimport math\nimport json\nimport time\nimport logging\nimport shutil\nimport numpy as np\nimport pandas as pd\nimport tensorflow as tf\nimport latticex.rosetta as rtt\nimport channel_sdk\n\n\nnp.set_printoptions(suppress=True)\ntf.compat.v1.logging.set_verbosity(tf.compat.v1.logging.ERROR)\nos.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'\nrtt.set_backend_loglevel(5)  # All(0), Trace(1), Debug(2), Info(3), Warn(4), Error(5), Fatal(6)\nlog = logging.getLogger(__name__)\n\nclass PrivacyLRTrain(object):\n    '''\n    Privacy logistic regression train base on rosetta.\n    '''\n\n    def __init__(self,\n                 channel_config: str,\n                 cfg_dict: dict,\n                 data_party: list,\n                 result_party: list,\n                 results_dir: str):\n        log.info(f\"channel_config:{channel_config}\")\n        log.info(f\"cfg_dict:{cfg_dict}\")\n        log.info(f\"data_party:{data_party}, result_party:{result_party}, results_dir:{results_dir}\")\n        assert isinstance(channel_config, str), \"type of channel_config must be str\"\n        assert isinstance(cfg_dict, dict), \"type of cfg_dict must be dict\"\n        assert isinstance(data_party, (list, tuple)), \"type of data_party must be list or tuple\"\n        assert isinstance(result_party, (list, tuple)), \"type of result_party must be list or tuple\"\n        assert isinstance(results_dir, str), \"type of results_dir must be str\"\n        \n        self.channel_config = channel_config\n        self.data_party = list(data_party)\n        self.result_party = list(result_party)\n        self.party_id = cfg_dict[\"party_id\"]\n        self.input_file = cfg_dict[\"data_party\"].get(\"input_file\")\n        self.key_column = cfg_dict[\"data_party\"].get(\"key_column\")\n        self.selected_columns = cfg_dict[\"data_party\"].get(\"selected_columns\")\n\n        dynamic_parameter = cfg_dict[\"dynamic_parameter\"]\n        self.label_owner = dynamic_parameter.get(\"label_owner\")\n        if self.party_id == self.label_owner:\n            self.label_column = dynamic_parameter.get(\"label_column\")\n            self.data_with_label = True\n        else:\n            self.label_column = \"\"\n            self.data_with_label = False\n                        \n        algorithm_parameter = dynamic_parameter[\"algorithm_parameter\"]\n        self.epochs = algorithm_parameter.get(\"epochs\", 10)\n        self.batch_size = algorithm_parameter.get(\"batch_size\", 256)\n        self.learning_rate = algorithm_parameter.get(\"learning_rate\", 0.001)\n        self.use_validation_set = algorithm_parameter.get(\"use_validation_set\", True)\n        self.validation_set_rate = algorithm_parameter.get(\"validation_set_rate\", 0.2)\n        self.predict_threshold = algorithm_parameter.get(\"predict_threshold\", 0.5)\n\n        self.output_file = os.path.join(results_dir, \"model\")\n        \n        self.check_parameters()\n\n    def check_parameters(self):\n        log.info(f\"check parameter start.\")        \n        assert self.epochs > 0, \"epochs must be greater 0\"\n        assert self.batch_size > 0, \"batch size must be greater 0\"\n        assert self.learning_rate > 0, \"learning rate must be greater 0\"\n        assert 0 < self.validation_set_rate < 1, \"validattion set rate must be between (0,1)\"\n        assert 0 <= self.predict_threshold <= 1, \"predict threshold must be between [0,1]\"\n        \n        if self.input_file:\n            self.input_file = self.input_file.strip()\n        if self.party_id in self.data_party:\n            if os.path.exists(self.input_file):\n                input_columns = pd.read_csv(self.input_file, nrows=0)\n                input_columns = list(input_columns.columns)\n                if self.key_column:\n                    assert self.key_column in input_columns, f\"key_column:{self.key_column} not in input_file\"\n                if self.selected_columns:\n                    error_col = []\n                    for col in self.selected_columns:\n                        if col not in input_columns:\n                            error_col.append(col)   \n                    assert not error_col, f\"selected_columns:{error_col} not in input_file\"\n                if self.label_column:\n                    assert self.label_column in input_columns, f\"label_column:{self.label_column} not in input_file\"\n            else:\n                raise Exception(f\"input_file is not exist. input_file={self.input_file}\")\n        log.info(f\"check parameter finish.\")\n                        \n        \n    def train(self):\n        '''\n        Logistic regression training algorithm implementation function\n        '''\n\n        log.info(\"extract feature or label.\")\n        train_x, train_y, val_x, val_y = self.extract_feature_or_label(with_label=self.data_with_label)\n        \n        log.info(\"start create and set channel.\")\n        self.create_set_channel()\n        log.info(\"waiting other party connect...\")\n        rtt.activate(\"SecureNN\")\n        log.info(\"protocol has been activated.\")\n        \n        log.info(f\"start set save model. save to party: {self.result_party}\")\n        rtt.set_saver_model(False, plain_model=self.result_party)\n        # sharing data\n        log.info(f\"start sharing train data. data_owner={self.data_party}, label_owner={self.label_owner}\")\n        shard_x, shard_y = rtt.PrivateDataset(data_owner=self.data_party, label_owner=self.label_owner).load_data(train_x, train_y, header=0)\n        log.info(\"finish sharing train data.\")\n        column_total_num = shard_x.shape[1]\n        log.info(f\"column_total_num = {column_total_num}.\")\n        \n        if self.use_validation_set:\n            log.info(\"start sharing validation data.\")\n            shard_x_val, shard_y_val = rtt.PrivateDataset(data_owner=self.data_party, label_owner=self.label_owner).load_data(val_x, val_y, header=0)\n            log.info(\"finish sharing validation data.\")\n\n        if self.party_id not in self.data_party:  \n            # mean the compute party and result party\n            log.info(\"compute start.\")\n            X = tf.placeholder(tf.float64, [None, column_total_num])\n            Y = tf.placeholder(tf.float64, [None, 1])\n            W = tf.Variable(tf.zeros([column_total_num, 1], dtype=tf.float64))\n            b = tf.Variable(tf.zeros([1], dtype=tf.float64))\n            logits = tf.matmul(X, W) + b\n            loss = tf.nn.sigmoid_cross_entropy_with_logits(labels=Y, logits=logits)\n            loss = tf.reduce_mean(loss)\n            # optimizer\n            optimizer = tf.train.GradientDescentOptimizer(self.learning_rate).minimize(loss)\n            init = tf.global_variables_initializer()\n            saver = tf.train.Saver(var_list=None, max_to_keep=5, name='v2')\n            \n            pred_Y = tf.sigmoid(tf.matmul(X, W) + b)\n            reveal_Y = rtt.SecureReveal(pred_Y)\n            actual_Y = tf.placeholder(tf.float64, [None, 1])\n            reveal_Y_actual = rtt.SecureReveal(actual_Y)\n\n            with tf.Session() as sess:\n                log.info(\"session init.\")\n                sess.run(init)\n                # train\n                log.info(\"train start.\")\n                train_start_time = time.time()\n                batch_num = math.ceil(len(shard_x) / self.batch_size)\n                for e in range(self.epochs):\n                    for i in range(batch_num):\n                        bX = shard_x[(i * self.batch_size): (i + 1) * self.batch_size]\n                        bY = shard_y[(i * self.batch_size): (i + 1) * self.batch_size]\n                        sess.run(optimizer, feed_dict={X: bX, Y: bY})\n                        if (i % 50 == 0) or (i == batch_num - 1):\n                            log.info(f\"epoch:{e + 1}/{self.epochs}, batch:{i + 1}/{batch_num}\")\n                log.info(f\"model save to: {self.output_file}\")\n                saver.save(sess, self.output_file)\n                train_use_time = round(time.time()-train_start_time, 3)\n                log.info(f\"save model success. train_use_time={train_use_time}s\")\n                \n                if self.use_validation_set:\n                    Y_pred = sess.run(reveal_Y, feed_dict={X: shard_x_val})\n                    log.debug(f\"Y_pred:\\n {Y_pred[:10]}\")\n                    Y_actual = sess.run(reveal_Y_actual, feed_dict={actual_Y: shard_y_val})\n                    log.debug(f\"Y_actual:\\n {Y_actual[:10]}\")\n        \n            running_stats = str(rtt.get_perf_stats(True)).replace('\\n', '').replace(' ', '')\n            log.info(f\"running stats: {running_stats}\")\n        else:\n            log.info(\"computing, please waiting for compute finish...\")\n        rtt.deactivate()\n     \n        log.info(\"remove temp dir.\")\n        if self.party_id in (self.data_party + self.result_party):\n            self.remove_temp_dir()\n        else:\n            # delete the model in the compute party.\n            self.remove_output_dir()\n        \n        if (self.party_id in self.result_party) and self.use_validation_set:\n            log.info(\"result_party evaluate model.\")\n            from sklearn.metrics import roc_auc_score, roc_curve, f1_score, precision_score, recall_score, accuracy_score\n            Y_pred_prob = Y_pred.astype(\"float\").reshape([-1, ])\n            Y_true = Y_actual.astype(\"float\").reshape([-1, ])\n            auc_score = roc_auc_score(Y_true, Y_pred_prob)\n            Y_pred_class = (Y_pred_prob > self.predict_threshold).astype('int64')  # default threshold=0.5\n            accuracy = accuracy_score(Y_true, Y_pred_class)\n            f1_score = f1_score(Y_true, Y_pred_class)\n            precision = precision_score(Y_true, Y_pred_class)\n            recall = recall_score(Y_true, Y_pred_class)\n            log.info(\"********************\")\n            log.info(f\"AUC: {round(auc_score, 6)}\")\n            log.info(f\"ACCURACY: {round(accuracy, 6)}\")\n            log.info(f\"F1_SCORE: {round(f1_score, 6)}\")\n            log.info(f\"PRECISION: {round(precision, 6)}\")\n            log.info(f\"RECALL: {round(recall, 6)}\")\n            log.info(\"********************\")\n        log.info(\"train finish.\")\n    \n    def create_set_channel(self):\n        '''\n        create and set channel.\n        '''\n        io_channel = channel_sdk.grpc.APIManager()\n        log.info(\"start create channel\")\n        channel = io_channel.create_channel(self.party_id, self.channel_config)\n        log.info(\"start set channel\")\n        rtt.set_channel(\"\", channel)\n        log.info(\"set channel success.\")\n    \n    def extract_feature_or_label(self, with_label: bool=False):\n        '''\n        Extract feature columns or label column from input file,\n        and then divide them into train set and validation set.\n        '''\n        train_x = \"\"\n        train_y = \"\"\n        val_x = \"\"\n        val_y = \"\"\n        temp_dir = self.get_temp_dir()\n        if self.party_id in self.data_party:\n            if self.input_file:\n                if with_label:\n                    usecols = self.selected_columns + [self.label_column]\n                else:\n                    usecols = self.selected_columns\n                \n                input_data = pd.read_csv(self.input_file, usecols=usecols, dtype=\"str\")\n                input_data = input_data[usecols]\n                # only if self.validation_set_rate==0, split_point==input_data.shape[0]\n                split_point = int(input_data.shape[0] * (1 - self.validation_set_rate))\n                assert split_point > 0, f\"train set is empty, because validation_set_rate:{self.validation_set_rate} is too big\"\n                \n                if with_label:\n                    y_data = input_data[self.label_column]\n                    train_y = os.path.join(temp_dir, f\"train_y_{self.party_id}.csv\")\n                    y_data.iloc[:split_point].to_csv(train_y, header=True, index=False)\n                    if self.use_validation_set:\n                        assert split_point < input_data.shape[0], \\\n                            f\"validation set is empty, because validation_set_rate:{self.validation_set_rate} is too small\"\n                        val_y = os.path.join(temp_dir, f\"val_y_{self.party_id}.csv\")\n                        y_data.iloc[split_point:].to_csv(val_y, header=True, index=False)\n                    del input_data[self.label_column]\n                \n                x_data = input_data\n                train_x = os.path.join(temp_dir, f\"train_x_{self.party_id}.csv\")\n                x_data.iloc[:split_point].to_csv(train_x, header=True, index=False)\n                if self.use_validation_set:\n                    assert split_point < input_data.shape[0], \\\n                            f\"validation set is empty, because validation_set_rate:{self.validation_set_rate} is too small\"\n                    val_x = os.path.join(temp_dir, f\"val_x_{self.party_id}.csv\")\n                    x_data.iloc[split_point:].to_csv(val_x, header=True, index=False)\n            else:\n                raise Exception(f\"data_node {self.party_id} not have data. input_file:{self.input_file}\")\n        return train_x, train_y, val_x, val_y\n    \n    def get_temp_dir(self):\n        '''\n        Get the directory for temporarily saving files\n        '''\n        temp_dir = os.path.join(os.path.dirname(self.output_file), 'temp')\n        if not os.path.exists(temp_dir):\n            os.makedirs(temp_dir, exist_ok=True)\n        return temp_dir\n\n    def remove_temp_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        Only delete temp file.\n        '''\n        temp_dir = self.get_temp_dir()\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n    \n    def remove_output_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        This is used to delete all output files of the non-resulting party\n        '''\n        temp_dir = os.path.dirname(self.output_file)\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n\n\ndef main(channel_config: str, cfg_dict: dict, data_party: list, result_party: list, results_dir: str):\n    '''\n    This is the entrance to this module\n    '''\n    privacy_lr = PrivacyLRTrain(channel_config, cfg_dict, data_party, result_party, results_dir)\n    privacy_lr.train()\n"
	//
	//predict_code := "# coding:utf-8\n\nimport sys\nsys.path.append(\"..\")\nimport os\nimport math\nimport json\nimport time\nimport logging\nimport shutil\nimport numpy as np\nimport pandas as pd\nimport tensorflow as tf\nimport latticex.rosetta as rtt\nimport channel_sdk\n\n\nnp.set_printoptions(suppress=True)\ntf.compat.v1.logging.set_verbosity(tf.compat.v1.logging.ERROR)\nos.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'\nrtt.set_backend_loglevel(5)  # All(0), Trace(1), Debug(2), Info(3), Warn(4), Error(5), Fatal(6)\nlog = logging.getLogger(__name__)\n\nclass PrivacyLRPredict(object):\n    '''\n    Privacy logistic regression predict base on rosetta.\n    '''\n\n    def __init__(self,\n                 channel_config: str,\n                 cfg_dict: dict,\n                 data_party: list,\n                 result_party: list,\n                 results_dir: str):\n        log.info(f\"channel_config:{channel_config}\")\n        log.info(f\"cfg_dict:{cfg_dict}\")\n        log.info(f\"data_party:{data_party}, result_party:{result_party}, results_dir:{results_dir}\")\n        assert isinstance(channel_config, str), \"type of channel_config must be str\"\n        assert isinstance(cfg_dict, dict), \"type of cfg_dict must be dict\"\n        assert isinstance(data_party, (list, tuple)), \"type of data_party must be list or tuple\"\n        assert isinstance(result_party, (list, tuple)), \"type of result_party must be list or tuple\"\n        assert isinstance(results_dir, str), \"type of results_dir must be str\"\n        \n        self.channel_config = channel_config\n        self.data_party = list(data_party)\n        self.result_party = list(result_party)\n        self.party_id = cfg_dict[\"party_id\"]\n        self.input_file = cfg_dict[\"data_party\"].get(\"input_file\")\n        self.key_column = cfg_dict[\"data_party\"].get(\"key_column\")\n        self.selected_columns = cfg_dict[\"data_party\"].get(\"selected_columns\")\n        dynamic_parameter = cfg_dict[\"dynamic_parameter\"]\n        self.model_restore_party = dynamic_parameter.get(\"model_restore_party\")\n        self.model_path = dynamic_parameter.get(\"model_path\")\n        self.model_file = os.path.join(self.model_path, \"model\")\n        self.predict_threshold = dynamic_parameter.get(\"predict_threshold\", 0.5)        \n        self.output_file = os.path.join(results_dir, \"result\")\n        self.data_party.remove(self.model_restore_party)  # except restore party\n        self.check_parameters()\n\n    def check_parameters(self):\n        log.info(f\"check parameter start.\")        \n        assert 0 <= self.predict_threshold <= 1, \"predict threshold must be between [0,1]\"\n        \n        if self.input_file:\n            self.input_file = self.input_file.strip()\n        if self.party_id in self.data_party:\n            if os.path.exists(self.input_file):\n                input_columns = pd.read_csv(self.input_file, nrows=0)\n                input_columns = list(input_columns.columns)\n                if self.key_column:\n                    assert self.key_column in input_columns, f\"key_column:{self.key_column} not in input_file\"\n                if self.selected_columns:\n                    error_col = []\n                    for col in self.selected_columns:\n                        if col not in input_columns:\n                            error_col.append(col)   \n                    assert not error_col, f\"selected_columns:{error_col} not in input_file\"\n            else:\n                raise Exception(f\"input_file is not exist. input_file={self.input_file}\")\n        if self.party_id == self.model_restore_party:\n            assert os.path.exists(self.model_path), f\"model path not found. model_path={self.model_path}\"\n        log.info(f\"check parameter finish.\")\n       \n\n    def predict(self):\n        '''\n        Logistic regression predict algorithm implementation function\n        '''\n\n        log.info(\"extract feature or id.\")\n        file_x, id_col = self.extract_feature_or_index()\n        \n        log.info(\"start create and set channel.\")\n        self.create_set_channel()\n        log.info(\"waiting other party connect...\")\n        rtt.activate(\"SecureNN\")\n        log.info(\"protocol has been activated.\")\n        \n        log.info(f\"start set restore model. restore party={self.model_restore_party}\")\n        rtt.set_restore_model(False, plain_model=self.model_restore_party)\n        # sharing data\n        log.info(f\"start sharing data. data_owner={self.data_party}\")\n        shard_x = rtt.PrivateDataset(data_owner=self.data_party).load_X(file_x, header=0)\n        log.info(\"finish sharing data .\")\n        column_total_num = shard_x.shape[1]\n        log.info(f\"column_total_num = {column_total_num}.\")\n\n        X = tf.placeholder(tf.float64, [None, column_total_num])\n        W = tf.Variable(tf.zeros([column_total_num, 1], dtype=tf.float64))\n        b = tf.Variable(tf.zeros([1], dtype=tf.float64))\n        saver = tf.train.Saver(var_list=None, max_to_keep=5, name='v2')\n        init = tf.global_variables_initializer()\n        # predict\n        pred_Y = tf.sigmoid(tf.matmul(X, W) + b)\n        reveal_Y = rtt.SecureReveal(pred_Y)  # only reveal to result party\n\n        with tf.Session() as sess:\n            log.info(\"session init.\")\n            sess.run(init)\n            log.info(\"start restore model.\")\n            if self.party_id == self.model_restore_party:\n                if os.path.exists(os.path.join(self.model_path, \"checkpoint\")):\n                    log.info(f\"model restore from: {self.model_file}.\")\n                    saver.restore(sess, self.model_file)\n                else:\n                    raise Exception(\"model not found or model damaged\")\n            else:\n                log.info(\"restore model...\")\n                temp_file = os.path.join(self.get_temp_dir(), 'ckpt_temp_file')\n                with open(temp_file, \"w\") as f:\n                    pass\n                saver.restore(sess, temp_file)\n            log.info(\"finish restore model.\")\n            \n            # predict\n            log.info(\"predict start.\")\n            predict_start_time = time.time()\n            Y_pred_prob = sess.run(reveal_Y, feed_dict={X: shard_x})\n            log.debug(f\"Y_pred_prob:\\n {Y_pred_prob[:10]}\")\n            predict_use_time = round(time.time() - predict_start_time, 3)\n            log.info(f\"predict success. predict_use_time={predict_use_time}s\")\n        rtt.deactivate()\n        log.info(\"rtt deactivate finish.\")\n        \n        if self.party_id in self.result_party:\n            log.info(\"predict result write to file.\")\n            output_file_predict_prob = os.path.splitext(self.output_file)[0] + \"_predict.csv\"\n            Y_pred_prob = Y_pred_prob.astype(\"float\")\n            Y_prob = pd.DataFrame(Y_pred_prob, columns=[\"Y_prob\"])\n            Y_class = (Y_pred_prob > self.predict_threshold) * 1\n            Y_class = pd.DataFrame(Y_class, columns=[\"Y_class\"])\n            Y_result = pd.concat([Y_prob, Y_class], axis=1)\n            Y_result.to_csv(output_file_predict_prob, header=True, index=False)\n        log.info(\"start remove temp dir.\")\n        self.remove_temp_dir()\n        log.info(\"predict finish.\")\n\n    def create_set_channel(self):\n        '''\n        create and set channel.\n        '''\n        io_channel = channel_sdk.grpc.APIManager()\n        log.info(\"start create channel\")\n        channel = io_channel.create_channel(self.party_id, self.channel_config)\n        log.info(\"start set channel\")\n        rtt.set_channel(\"\", channel)\n        log.info(\"set channel success.\")\n        \n    def extract_feature_or_index(self):\n        '''\n        Extract feature columns or index column from input file.\n        '''\n        file_x = \"\"\n        id_col = None\n        temp_dir = self.get_temp_dir()\n        if self.party_id in self.data_party:\n            if self.input_file:\n                usecols = [self.key_column] + self.selected_columns\n                input_data = pd.read_csv(self.input_file, usecols=usecols, dtype=\"str\")\n                input_data = input_data[usecols]\n                id_col = input_data[self.key_column]\n                file_x = os.path.join(temp_dir, f\"file_x_{self.party_id}.csv\")\n                x_data = input_data.drop(labels=self.key_column, axis=1)\n                x_data.to_csv(file_x, header=True, index=False)\n            else:\n                raise Exception(f\"data_party:{self.party_id} not have data. input_file:{self.input_file}\")\n        return file_x, id_col\n    \n    def get_temp_dir(self):\n        '''\n        Get the directory for temporarily saving files\n        '''\n        temp_dir = os.path.join(os.path.dirname(self.output_file), 'temp')\n        if not os.path.exists(temp_dir):\n            os.makedirs(temp_dir, exist_ok=True)\n        return temp_dir\n\n    def remove_temp_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        Only delete temp file.\n        '''\n        temp_dir = self.get_temp_dir()\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n\n\ndef main(channel_config: str, cfg_dict: dict, data_party: list, result_party: list, results_dir: str):\n    '''\n    This is the entrance to this module\n    '''\n    privacy_lr = PrivacyLRPredict(channel_config, cfg_dict, data_party, result_party, results_dir)\n    privacy_lr.predict()\n"
	//
	//train_params := "{\n    \"label_owner\": \"p2\",       \n    \"label_column\": \"Y\",      \n    \"algorithm_parameter\": {   \n      \"epochs\": 10,           \n      \"batch_size\": 256,      \n      \"learning_rate\": 0.1,    \n      \"use_validation_set\": true,  \n      \"validation_set_rate\": 0.2,  \n      \"predict_threshold\": 0.5  \n    }\n}"
	//
	////predict_params := "{\n    \"model_restore_party\": \"p2\", \n    \"model_path\": \"/home/user1/fighter/results30001/task:0x292f8d472b51578012e753acb56182601b78d6d5c8c94170e186448253ca0d58/p7\",    \n    \"predict_threshold\": 0.5   \n}"
	//
	//predict_params := "{\n    \"model_restore_party\": \"p2\", \n    \"model_path\": \"/home/user1/fighter/results30001/task:0x2062cbde963d2c2efda345dbd1ad080ca8a85eabb01e2738c7ae8a9bb9889875/p7\",    \n    \"predict_threshold\": 0.5   \n}"

	//arr :=  [][]int{[]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{3, 4}, []int{1, 2}}
	//loop:  // 1, 2, 2, 1, 2
	//for _, arr := range [][]int{[]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{3, 4}, []int{1, 2}} {
	//	for _, v := range arr {
	//		if v == 3 {
	//			continue loop
	//		} else {
	//			fmt.Println(v)
	//		}
	//	}
	//}

	//str := "jobNode:0x0375653f3c2871884eb2a7ed3c1ef61d3d6e2bc31ddd0a80f982fb48fc4d6073"
	//fmt.Println(len([]byte(str)))
	//
	//partyId := "y2"
	//partyIds := []string{"y1", "y2", "y3"}
	//for i, id := range partyIds {
	//	if id == partyId {
	//		partyIds = append(partyIds[:i], partyIds[i+1:]...)
	//	}
	//}
	//if len(partyIds) == 0 {
	//	fmt.Println("empty", partyIds)
	//} else {
	//	fmt.Println("non-empty", partyIds)
	//}
	//
	//
	//code := "# coding:utf-8\n\nimport os\nimport sys\nimport math\nimport json\nimport time\nimport logging\nimport shutil\nimport numpy as np\nimport pandas as pd\nimport tensorflow as tf\nimport latticex.rosetta as rtt\nimport channel_sdk\n\n\nnp.set_printoptions(suppress=True)\ntf.compat.v1.logging.set_verbosity(tf.compat.v1.logging.ERROR)\nos.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'\nrtt.set_backend_loglevel(5)  # All(0), Trace(1), Debug(2), Info(3), Warn(4), Error(5), Fatal(6)\nlog = logging.getLogger(__name__)\n\nclass PrivacyLRTrain(object):\n    '''\n    Privacy logistic regression train base on rosetta.\n    '''\n\n    def __init__(self,\n                 channel_config: str,\n                 cfg_dict: dict,\n                 data_party: list,\n                 result_party: list,\n                 results_dir: str):\n        log.info(f\"channel_config:{channel_config}\")\n        log.info(f\"cfg_dict:{cfg_dict}\")\n        log.info(f\"data_party:{data_party}, result_party:{result_party}, results_dir:{results_dir}\")\n        assert isinstance(channel_config, str), \"type of channel_config must be str\"\n        assert isinstance(cfg_dict, dict), \"type of cfg_dict must be dict\"\n        assert isinstance(data_party, (list, tuple)), \"type of data_party must be list or tuple\"\n        assert isinstance(result_party, (list, tuple)), \"type of result_party must be list or tuple\"\n        assert isinstance(results_dir, str), \"type of results_dir must be str\"\n        \n        self.channel_config = channel_config\n        self.data_party = list(data_party)\n        self.result_party = list(result_party)\n        self.party_id = cfg_dict[\"party_id\"]\n        self.input_file = cfg_dict[\"data_party\"].get(\"input_file\")\n        self.key_column = cfg_dict[\"data_party\"].get(\"key_column\")\n        self.selected_columns = cfg_dict[\"data_party\"].get(\"selected_columns\")\n\n        dynamic_parameter = cfg_dict[\"dynamic_parameter\"]\n        self.label_owner = dynamic_parameter.get(\"label_owner\")\n        if self.party_id == self.label_owner:\n            self.label_column = dynamic_parameter.get(\"label_column\")\n            self.data_with_label = True\n        else:\n            self.label_column = \"\"\n            self.data_with_label = False\n                        \n        algorithm_parameter = dynamic_parameter[\"algorithm_parameter\"]\n        self.epochs = algorithm_parameter.get(\"epochs\", 10)\n        self.batch_size = algorithm_parameter.get(\"batch_size\", 256)\n        self.learning_rate = algorithm_parameter.get(\"learning_rate\", 0.001)\n        self.use_validation_set = algorithm_parameter.get(\"use_validation_set\", True)\n        self.validation_set_rate = algorithm_parameter.get(\"validation_set_rate\", 0.2)\n        self.predict_threshold = algorithm_parameter.get(\"predict_threshold\", 0.5)\n\n        self.output_file = os.path.join(results_dir, \"model\")\n        \n        self.check_parameters()\n\n    def check_parameters(self):\n        log.info(f\"check parameter start.\")        \n        assert isinstance(self.epochs, int) and self.epochs > 0, \"epochs must be type(int) and greater 0\"\n        assert isinstance(self.batch_size, int) and self.batch_size > 0, \"batch_size must be type(int) and greater 0\"\n        assert isinstance(self.learning_rate, float) and self.learning_rate > 0, \"learning rate must be type(float) and greater 0\"\n        assert 0 < self.validation_set_rate < 1, \"validattion set rate must be between (0,1)\"\n        assert 0 <= self.predict_threshold <= 1, \"predict threshold must be between [0,1]\"\n        \n        if self.input_file:\n            self.input_file = self.input_file.strip()\n        if self.party_id in self.data_party:\n            if os.path.exists(self.input_file):\n                input_columns = pd.read_csv(self.input_file, nrows=0)\n                input_columns = list(input_columns.columns)\n                if self.key_column:\n                    assert self.key_column in input_columns, f\"key_column:{self.key_column} not in input_file\"\n                if self.selected_columns:\n                    error_col = []\n                    for col in self.selected_columns:\n                        if col not in input_columns:\n                            error_col.append(col)   \n                    assert not error_col, f\"selected_columns:{error_col} not in input_file\"\n                if self.label_column:\n                    assert self.label_column in input_columns, f\"label_column:{self.label_column} not in input_file\"\n            else:\n                raise Exception(f\"input_file is not exist. input_file={self.input_file}\")\n        log.info(f\"check parameter finish.\")\n                        \n        \n    def train(self):\n        '''\n        Logistic regression training algorithm implementation function\n        '''\n\n        log.info(\"extract feature or label.\")\n        train_x, train_y, val_x, val_y = self.extract_feature_or_label(with_label=self.data_with_label)\n        \n        log.info(\"start create and set channel.\")\n        self.create_set_channel()\n        log.info(\"waiting other party connect...\")\n        rtt.activate(\"SecureNN\")\n        log.info(\"protocol has been activated.\")\n        \n        log.info(f\"start set save model. save to party: {self.result_party}\")\n        rtt.set_saver_model(False, plain_model=self.result_party)\n        # sharing data\n        log.info(f\"start sharing train data. data_owner={self.data_party}, label_owner={self.label_owner}\")\n        shard_x, shard_y = rtt.PrivateDataset(data_owner=self.data_party, label_owner=self.label_owner).load_data(train_x, train_y, header=0)\n        log.info(\"finish sharing train data.\")\n        column_total_num = shard_x.shape[1]\n        log.info(f\"column_total_num = {column_total_num}.\")\n        \n        if self.use_validation_set:\n            log.info(\"start sharing validation data.\")\n            shard_x_val, shard_y_val = rtt.PrivateDataset(data_owner=self.data_party, label_owner=self.label_owner).load_data(val_x, val_y, header=0)\n            log.info(\"finish sharing validation data.\")\n\n        if self.party_id not in self.data_party:  \n            # mean the compute party and result party\n            log.info(\"compute start.\")\n            X = tf.placeholder(tf.float64, [None, column_total_num])\n            Y = tf.placeholder(tf.float64, [None, 1])\n            W = tf.Variable(tf.zeros([column_total_num, 1], dtype=tf.float64))\n            b = tf.Variable(tf.zeros([1], dtype=tf.float64))\n            logits = tf.matmul(X, W) + b\n            loss = tf.nn.sigmoid_cross_entropy_with_logits(labels=Y, logits=logits)\n            loss = tf.reduce_mean(loss)\n            # optimizer\n            optimizer = tf.train.GradientDescentOptimizer(self.learning_rate).minimize(loss)\n            init = tf.global_variables_initializer()\n            saver = tf.train.Saver(var_list=None, max_to_keep=5, name='v2')\n            \n            pred_Y = tf.sigmoid(tf.matmul(X, W) + b)\n            reveal_Y = rtt.SecureReveal(pred_Y)\n            actual_Y = tf.placeholder(tf.float64, [None, 1])\n            reveal_Y_actual = rtt.SecureReveal(actual_Y)\n\n            with tf.Session() as sess:\n                log.info(\"session init.\")\n                sess.run(init)\n                # train\n                log.info(\"train start.\")\n                train_start_time = time.time()\n                batch_num = math.ceil(len(shard_x) / self.batch_size)\n                for e in range(self.epochs):\n                    for i in range(batch_num):\n                        bX = shard_x[(i * self.batch_size): (i + 1) * self.batch_size]\n                        bY = shard_y[(i * self.batch_size): (i + 1) * self.batch_size]\n                        sess.run(optimizer, feed_dict={X: bX, Y: bY})\n                        if (i % 50 == 0) or (i == batch_num - 1):\n                            log.info(f\"epoch:{e + 1}/{self.epochs}, batch:{i + 1}/{batch_num}\")\n                log.info(f\"model save to: {self.output_file}\")\n                saver.save(sess, self.output_file)\n                train_use_time = round(time.time()-train_start_time, 3)\n                log.info(f\"save model success. train_use_time={train_use_time}s\")\n                \n                if self.use_validation_set:\n                    Y_pred = sess.run(reveal_Y, feed_dict={X: shard_x_val})\n                    log.info(f\"Y_pred:\\n {Y_pred[:10]}\")\n                    Y_actual = sess.run(reveal_Y_actual, feed_dict={actual_Y: shard_y_val})\n                    log.info(f\"Y_actual:\\n {Y_actual[:10]}\")\n        \n            running_stats = str(rtt.get_perf_stats(True)).replace('\\n', '').replace(' ', '')\n            log.info(f\"running stats: {running_stats}\")\n        else:\n            log.info(\"computing, please waiting for compute finish...\")\n        rtt.deactivate()\n     \n        log.info(\"remove temp dir.\")\n        if self.party_id in (self.data_party + self.result_party):\n            # self.remove_temp_dir()\n            pass\n        else:\n            # delete the model in the compute party.\n            self.remove_output_dir()\n        \n        if (self.party_id in self.result_party) and self.use_validation_set:\n            log.info(\"result_party evaluate model.\")\n            from sklearn.metrics import roc_auc_score, roc_curve, f1_score, precision_score, recall_score, accuracy_score\n            Y_pred_prob = Y_pred.astype(\"float\").reshape([-1, ])\n            Y_true = Y_actual.astype(\"float\").reshape([-1, ])\n            auc_score = roc_auc_score(Y_true, Y_pred_prob)\n            log.info(f\"AUC: {round(auc_score, 6)}\")\n            Y_pred_class = (Y_pred_prob > self.predict_threshold).astype('int64')  # default threshold=0.5\n            accuracy = accuracy_score(Y_true, Y_pred_class)\n            log.info(f\"ACCURACY: {round(accuracy, 6)}\")\n            f1_score = f1_score(Y_true, Y_pred_class)\n            precision = precision_score(Y_true, Y_pred_class)\n            recall = recall_score(Y_true, Y_pred_class)\n            log.info(\"********************\")\n            log.info(f\"AUC: {round(auc_score, 6)}\")\n            log.info(f\"ACCURACY: {round(accuracy, 6)}\")\n            log.info(f\"F1_SCORE: {round(f1_score, 6)}\")\n            log.info(f\"PRECISION: {round(precision, 6)}\")\n            log.info(f\"RECALL: {round(recall, 6)}\")\n            log.info(\"********************\")\n        log.info(\"train finish.\")\n    \n    def create_set_channel(self):\n        '''\n        create and set channel.\n        '''\n        io_channel = channel_sdk.grpc.APIManager()\n        log.info(\"start create channel\")\n        channel = io_channel.create_channel(self.party_id, self.channel_config)\n        log.info(\"start set channel\")\n        rtt.set_channel(\"\", channel)\n        log.info(\"set channel success.\")\n    \n    def extract_feature_or_label(self, with_label: bool=False):\n        '''\n        Extract feature columns or label column from input file,\n        and then divide them into train set and validation set.\n        '''\n        train_x = \"\"\n        train_y = \"\"\n        val_x = \"\"\n        val_y = \"\"\n        temp_dir = self.get_temp_dir()\n        if self.party_id in self.data_party:\n            if self.input_file:\n                if with_label:\n                    usecols = self.selected_columns + [self.label_column]\n                else:\n                    usecols = self.selected_columns\n                \n                input_data = pd.read_csv(self.input_file, usecols=usecols, dtype=\"str\")\n                input_data = input_data[usecols]\n                # only if self.validation_set_rate==0, split_point==input_data.shape[0]\n                split_point = int(input_data.shape[0] * (1 - self.validation_set_rate))\n                assert split_point > 0, f\"train set is empty, because validation_set_rate:{self.validation_set_rate} is too big\"\n                \n                if with_label:\n                    y_data = input_data[self.label_column]\n                    train_y_data = y_data.iloc[:split_point]\n                    train_class_num = train_y_data.unique().shape[0]\n                    assert train_class_num == 2, f\"train set must be 2 class, not {train_class_num} class.\"\n                    train_y = os.path.join(temp_dir, f\"train_y_{self.party_id}.csv\")\n                    train_y_data.to_csv(train_y, header=True, index=False)\n                    if self.use_validation_set:\n                        assert split_point < input_data.shape[0], \\\n                            f\"validation set is empty, because validation_set_rate:{self.validation_set_rate} is too small\"\n                        val_y_data = y_data.iloc[split_point:]\n                        val_class_num = val_y_data.unique().shape[0]\n                        assert val_class_num == 2, f\"validation set must be 2 class, not {val_class_num} class.\"\n                        val_y = os.path.join(temp_dir, f\"val_y_{self.party_id}.csv\")\n                        val_y_data.to_csv(val_y, header=True, index=False)\n                    del input_data[self.label_column]\n                \n                x_data = input_data\n                train_x = os.path.join(temp_dir, f\"train_x_{self.party_id}.csv\")\n                x_data.iloc[:split_point].to_csv(train_x, header=True, index=False)\n                if self.use_validation_set:\n                    assert split_point < input_data.shape[0], \\\n                            f\"validation set is empty, because validation_set_rate:{self.validation_set_rate} is too small.\"\n                    val_x = os.path.join(temp_dir, f\"val_x_{self.party_id}.csv\")\n                    x_data.iloc[split_point:].to_csv(val_x, header=True, index=False)\n            else:\n                raise Exception(f\"data_node {self.party_id} not have data. input_file:{self.input_file}\")\n        return train_x, train_y, val_x, val_y\n    \n    def get_temp_dir(self):\n        '''\n        Get the directory for temporarily saving files\n        '''\n        temp_dir = os.path.join(os.path.dirname(self.output_file), 'temp')\n        if not os.path.exists(temp_dir):\n            os.makedirs(temp_dir, exist_ok=True)\n        return temp_dir\n\n    def remove_temp_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        Only delete temp file.\n        '''\n        temp_dir = self.get_temp_dir()\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n    \n    def remove_output_dir(self):\n        '''\n        Delete all files in the temporary directory, these files are some temporary data.\n        This is used to delete all output files of the non-resulting party\n        '''\n        temp_dir = os.path.dirname(self.output_file)\n        if os.path.exists(temp_dir):\n            shutil.rmtree(temp_dir)\n\n\ndef main(channel_config: str, cfg_dict: dict, data_party: list, result_party: list, results_dir: str):\n    '''\n    This is the entrance to this module\n    '''\n    privacy_lr = PrivacyLRTrain(channel_config, cfg_dict, data_party, result_party, results_dir)\n    privacy_lr.train()\n"

	//x := new(big.Int).Div(new(big.Int).SetUint64(8), new(big.Int).SetUint64(2097152))
	//y := new(big.Int).Mod(new(big.Int).SetUint64(8), new(big.Int).SetUint64(2097152))

	//  {"mem": 2097152, "processor": 1, "bandwidth": 65536}, cost.mem: {1073741824}, cost.Bandwidth: {3145728}, cost.Processor: {1}, return needSlotCount: {512}

	//fmt.Println(DivCeil(2097153, 2097152))
	//fmt.Println(DivCeil(1073741824, 2097152), 1024 * 1024 * 1024 * 1)    	// 1.00 GB  1073741824   mem 512 slot     512 slot,    4 slot
	//fmt.Println(DivCeil(3145728, 65536), 1024 * 1024 * 4)         		// 4.00 MBP/S            band  48 slot
	//fmt.Println(DivCeil(1, 1), 20000000/4194304, 1024 * 64)                          // 1 slot
	//fmt.Println(false && true)
	//20000000  == 16mbps
	//4194304
	//decimal.NewFromInt(int64(2097152))
	//fmt.Printf("x: %d, y : %d, z: %d \n", x, y, 8 / 2097152)

	//k, err := crypto.GenerateKey()
	//if nil != err {
	//	fmt.Println("Failed to generate random NodeId private key", err)
	//	return
	//}
	//data := []byte("identity_979d91441f904f08b5d74814d8d30e9b")
	//proof, err := vrf.Prove(k, data)
	//if nil != err {
	//	fmt.Println("Failed to generate vrf proof", err)
	//	return
	//}
	//
	//hash := vrf.ProofToHash(proof)
	//fmt.Println("proof hash len", len(hash), "hash", common.BytesToHash(hash).String())
	//
	//flag, err := vrf.Verify(&(k.PublicKey), proof, data)
	//
	//if nil != err {
	//	fmt.Println("Failed to verify vrf proof", err)
	//	return
	//}
	//
	//fmt.Println("Verify result", flag)

	//a := 10_000
	//b := 10000
	//fmt.Println(a, b, a == b)

	//start := time.Now()
	//time.Sleep(10 * time.Second)
	//fmt.Printf("duration: %d \n", time.Since(start).Milliseconds())

	//fmt.Println(math.MaxInt32, 1 << 23, 1024*1024*8, 1 << 22, 1024*1024*4)
	//
	//
	//fmt.Println(len("cd99cca4de60c91585c2ebc1c54b95b91bed30c2455c3dce97d4945a8501cbf1efc1ac4d98812fba55506b568d1affcf741d706eda26d99d6d94dd2e182d379b"))

	//signal := make(chan chan string)
	//go func(signal chan chan string) {
	//	time.Sleep(4*time.Second)
	//	<- signal <- "hello"
	//}(signal)
	//c := make(chan string)
	//signal <- c
	//fmt.Println(<- c)

	//ctx, cancelFn := context.WithCancel(context.Background())
	//
	//start := time.Now()
	//go func(cancelFn context.CancelFunc) {
	//
	//	time.Sleep(2*time.Second)
	//	cancelFn()
	//
	//}(cancelFn)
	//<-ctx.Done()
	//fmt.Println("duration:", time.Since(start))

	/*now := time.Now()

	fmt.Println("now time ", now.Format("2006-01-02 15:04:05"))

	queue := make(TaskBullets, 0)

	heap.Push(&queue, &TaskBullet{
		Name: "a",
		Prioty: int64(now.Add(time.Duration(6) * time.Second).Unix()),
	})
	heap.Push(&queue, &TaskBullet{
		Name: "b",
		Prioty: int64(now.Add(time.Duration(3) * time.Second).Unix()),
	})
	heap.Push(&queue, &TaskBullet{
		Name: "c",
		Prioty: int64(now.Add(time.Duration(8) * time.Second).Unix()),
	})


	ctx, cancelFn := context.WithCancel(context.Background())


	go func(cancelFn context.CancelFunc, queue *TaskBullets) {

		fmt.Println("Start handle queue")

		timer := time.NewTimer(time.Unix(queue.TimeSleepUntil(), 0).Sub(time.Now()))

		go func() {
			fmt.Println("Start add new one member into queue")
			x := &TaskBullet{
				Name: "d",
				Prioty: int64(now.Add(time.Duration(1) * time.Second).Unix()),
			}
			heap.Push(queue, x)
			timer.Reset(time.Unix(x.Prioty, 0).Sub(time.Now()))
		}()

		for {
			select {
			case <- timer.C:


				m := heap.Pop(queue)
				x := m.(*TaskBullet)

				pri:
				fmt.Printf("I am %s, prioty %s \n", x.Name, time.Unix(x.Prioty, 0).Format("2006-01-02 15:04:05"))

				if len(*queue) == 0 {
					cancelFn()
					return
				}

				for {
					if queue.TimeSleepUntil() == x.Prioty || queue.TimeSleepUntil() < x.Prioty {
						m = heap.Pop(queue)
						x = m.(*TaskBullet)
						goto pri
					} else {
						timer.Reset(time.Unix(queue.TimeSleepUntil(), 0).Sub(time.Now()))
						break
					}
				}
			}
		}

	}(cancelFn, &queue)
	<-ctx.Done()
	fmt.Println("duration:", time.Since(now))

	for {
		var x *int32   	// 因为 x 此时 指向的是 nil
		*x = 0			// 所以,  解引用 *nil 肯定导致 崩溃啊
	}*/

	//now := timeutils.UnixMsec()
	//fmt.Println(now)

	now := time.Now()
	after1 := now.Add(time.Duration(6) * time.Second)
	after2 := now.Add(time.Duration(1) * time.Second)
	after3 := now.Add(time.Duration(8) * time.Second)
	after4 := now.Add(time.Duration(4) * time.Second)
	after5 := now.Add(time.Duration(2) * time.Second)
	after6 := now.Add(time.Duration(3) * time.Second)
	after7 := now.Add(time.Duration(1) * time.Second)

	queue := NewSyncExecuteTaskMonitorQueue(0)

	fmt.Println("now time ", now.Format("2006-01-02 15:04:05"), "timestamp", now.UnixNano()/1e6)

	timer := queue.Timer()
	timer.Reset(time.Duration(math.MaxInt32) * time.Millisecond)

	ctx, cancelFn := context.WithCancel(context.Background())

	go func(cancelFn context.CancelFunc, queue *SyncExecuteTaskMonitorQueue) {

		fmt.Println("Start handle queue")

		//timer := time.NewTimer(time.Duration(queue.TimeSleepUntil() - timeutils.UnixMsec()) * time.Millisecond)

		for {
			select {
			case <-timer.C:

				then := time.Now().UnixNano() / 1e6

				queue.lock.Lock()

			rerun:
				for len(*(queue.queue)) > 0 {
					if future := queue.RunMonitor(then); future != 0 {
						if future > 0 {
							then = timeutils.UnixMsec()
							if future > then {
								timer.Reset(time.Duration(future-then) * time.Millisecond)
								break
							} else {
								continue rerun
							}
						}
					}
				}

				queue.lock.Unlock()

				if len(*(queue.queue)) == 0 {
					cancelFn()
					return
				}
				//cancelFn()
				//return
			}
		}

	}(cancelFn, queue)

	go func(queue *SyncExecuteTaskMonitorQueue) {
		fmt.Println("Start add new one member into queue")

		queue.AddMonitor(NewExecuteTaskMonitor("A", after1.UnixNano()/1e6, func() {
			fmt.Println("Removed A, after1 time ", after1.Format("2006-01-02 15:04:05"), "timestamp", after1.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		queue.AddMonitor(NewExecuteTaskMonitor("B", after2.UnixNano()/1e6, func() {
			fmt.Println("Removed B, after2 time ", after2.Format("2006-01-02 15:04:05"), "timestamp", after2.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		queue.AddMonitor(NewExecuteTaskMonitor("C", after3.UnixNano()/1e6, func() {
			fmt.Println("Removed C, after3 time ", after3.Format("2006-01-02 15:04:05"), "timestamp", after3.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		queue.AddMonitor(NewExecuteTaskMonitor("D", after4.UnixNano()/1e6, func() {
			fmt.Println("Removed D, after4 time ", after4.Format("2006-01-02 15:04:05"), "timestamp", after4.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		queue.AddMonitor(NewExecuteTaskMonitor("E", after5.UnixNano()/1e6, func() {
			fmt.Println("Removed E, after5 time ", after5.Format("2006-01-02 15:04:05"), "timestamp", after5.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		queue.AddMonitor(NewExecuteTaskMonitor("F", after6.UnixNano()/1e6, func() {
			fmt.Println("Removed F, after6 time ", after6.Format("2006-01-02 15:04:05"), "timestamp", after6.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		queue.AddMonitor(NewExecuteTaskMonitor("G", after7.UnixNano()/1e6, func() {
			fmt.Println("Removed G, after7 time ", after7.Format("2006-01-02 15:04:05"), "timestamp", after7.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
		}))

		for i := 0; i < 2; i++ {
			//var du int64
			//du = int64(i % 3)
			//if du == 0 {
			//	du = 1
			//}
			//then := time.Now().Add(time.Duration(du) * time.Second)
			//time.Sleep(time.Duration(du-1) * time.Second)
			//name := fmt.Sprintf("com_%d", i)
			//queue.AddMonitor(NewExecuteTaskMonitor(name, then.UnixNano()/1e6, func() {
			//	fmt.Println(fmt.Sprintf("Removed %s, then time ", name), then.Format("2006-01-02 15:04:05"), "timestamp", then.UnixNano()/1e6, "now", time.Now().Format("2006-01-02 15:04:05"))
			//}))
			queue.DelMonitor(i)
		}

	}(queue)

	<-ctx.Done()
	fmt.Println("duration:", time.Since(now))

	////////
	//future := time.Duration(-1 - timeutils.UnixMsec())
	//fmt.Println(future, future < 0)
	//if future <= 0 {
	//	future = 0
	//}
	//fmt.Println("now:", time.Now().Format("2006-01-02 15:04:05"))
	//taskMonitorTicker := time.NewTimer(future * time.Millisecond)
	//<- taskMonitorTicker.C
	//fmt.Println("then:", time.Now().Format("2006-01-02 15:04:05"))

	timestamp := time.Now().UnixNano() / 1e6

	datetime := time.Unix(timestamp/1000, 0).Format("2006-01-02 15:04:05")
	fmt.Println(datetime)
}

//type TaskBullet struct {
//	Name   string
//	Prioty int64
//}
//type TaskBullets []*TaskBullet
//
//func (h TaskBullets) Len() int           { return len(h) }
//func (h TaskBullets) Less(i, j int) bool { return h[i].Prioty < h[j].Prioty } // term:  a.3 < c.2 < b.1,  So order is: a c b
//func (h TaskBullets) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *TaskBullets) Push(x interface{}) {
//	m := x.(*TaskBullet)
//	fmt.Printf("Push mem %s, target time %s \n", m.Name, time.Unix(m.Prioty, 0).Format("2006-01-02 15:04:05"))
//	*h = append(*h, m)
//}
//
//func (h *TaskBullets) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	//fmt.Printf("Pop mem %s, target time %s \n", x.Name, time.Unix(x.Prioty, 0).Format("2006-01-02 15:04:05"))
//	return x
//}
//
//func (h *TaskBullets) TimeSleepUntil() int64 {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	return x.Prioty
//}
//
//func DivCeil(a, b uint64) uint64 {
//	div := a / b
//	mod := a % b
//
//	if mod > 0 {
//		div += 1
//	}
//	return div
//}
//
//func fA() (r int) {
//	t := 5
//	defer func() {
//		t = t + 5
//	}()
//	return t
//}
//
//func fB() int {
//	t := 5
//	defer func() {
//		t = t + 5
//	}()
//	return t
//}
//
//func fC() (r int) {
//	defer func(r int) {
//		r = r + 5
//	}(r)
//	return 1
//}

//

type ExecuteTaskMonitor struct {
	taskId string
	index  int
	when   int64 // target timestamp
	fn     func()
}

func NewExecuteTaskMonitor(taskId string, when int64, fn func()) *ExecuteTaskMonitor {
	fmt.Printf("New a monitor, taskId: %s\n", taskId)
	return &ExecuteTaskMonitor{
		taskId:  taskId,
		when:    when,
		fn:      fn,
	}
}

func (ett *ExecuteTaskMonitor) String() string  {
	return fmt.Sprintf(`{"index": %d, "taskId": "%s", "when": %d}`, ett.index, ett.taskId, ett.when)
}
func (ett *ExecuteTaskMonitor) GetTaskId() string  { return ett.taskId }
func (ett *ExecuteTaskMonitor) GetIndex() int { return ett.index }
func (ett *ExecuteTaskMonitor) GetWhen() int64     { return ett.when }

type executeTaskMonitorQueue []*ExecuteTaskMonitor
func (queue *executeTaskMonitorQueue) String() string {
	arr := make([]string, len(*queue))
	for i, ett := range *queue {
		arr[i] = ett.String()
	}
	return "[" + strings.Join(arr, ",") +  "]"
}

type SyncExecuteTaskMonitorQueue struct {
	lock  sync.Mutex
	timer *time.Timer
	queue *executeTaskMonitorQueue
}

func NewSyncExecuteTaskMonitorQueue(size int) *SyncExecuteTaskMonitorQueue {
	queue := make(executeTaskMonitorQueue, size)
	timer := time.NewTimer(0)
	<-timer.C
	return &SyncExecuteTaskMonitorQueue{
		queue: &(queue),
		timer: timer,
	}
}

func (syncQueue *SyncExecuteTaskMonitorQueue) printQueue() {
	 fmt.Printf("Print queue ==> %s\n", syncQueue.queue.String())
}

func (syncQueue *SyncExecuteTaskMonitorQueue) Timer() *time.Timer {
	return syncQueue.timer
}

func (syncQueue *SyncExecuteTaskMonitorQueue) AddMonitor(m *ExecuteTaskMonitor) {
	syncQueue.lock.Lock()
	defer syncQueue.lock.Unlock()
	// when must never be negative;
	if m.when-timeutils.UnixMsec() < 0 {
		panic("target time is negative number")
	}
	i := len(*(syncQueue.queue))
	m.index = i
	*(syncQueue.queue) = append(*(syncQueue.queue), m)
	syncQueue.siftUpMonitor(i)

	// reset the timer
	var until int64
	if len(*(syncQueue.queue)) > 0 {
		until = (*(syncQueue.queue))[0].when
	} else {
		until = -1
	}
	future := time.Duration(until - timeutils.UnixMsec())
	if future <= 0 {
		future = 0
	}
	syncQueue.timer.Reset(future * time.Millisecond)
	syncQueue.printQueue()
}

func (syncQueue *SyncExecuteTaskMonitorQueue) DelMonitor(i int) {

	syncQueue.lock.Lock()
	defer syncQueue.lock.Unlock()

	fmt.Printf("Deleted %s\n", (*(syncQueue.queue))[i].GetTaskId())

	last := len(*(syncQueue.queue)) - 1
	if i != last {
		(*(syncQueue.queue))[last].index = i
		(*(syncQueue.queue))[i] = (*(syncQueue.queue))[last]
	}
	(*(syncQueue.queue))[last] = nil
	*(syncQueue.queue) = (*(syncQueue.queue))[:last]
	if i != last {
		// Moving to i may have moved the last monitor to a new parent,
		// so sift up to preserve the heap guarantee.
		syncQueue.siftUpMonitor(i)
		syncQueue.siftDownMonitor(i)
	}
	syncQueue.printQueue()
}

func (syncQueue *SyncExecuteTaskMonitorQueue) delMonitor0() {

	last := len(*(syncQueue.queue)) - 1
	if last > 0 {
		(*(syncQueue.queue))[last].index = 0
		(*(syncQueue.queue))[0] = (*(syncQueue.queue))[last]
	}
	(*(syncQueue.queue))[last] = nil
	*(syncQueue.queue) = (*(syncQueue.queue))[:last]

	if last > 0 {
		syncQueue.siftDownMonitor(0)
	}
	syncQueue.printQueue()
}

func (syncQueue *SyncExecuteTaskMonitorQueue) RunMonitor(now int64) int64 {
	if len(*(syncQueue.queue)) == 0 {
		return 0
	}

	m := (*(syncQueue.queue))[0]
	if m.when > now {
		// Not ready to run.
		return m.when
	}
	f := m.fn
	// Remove from heap.
	syncQueue.delMonitor0()
	syncQueue.lock.Unlock()
	f()
	syncQueue.lock.Lock()
	return 0
}

func (syncQueue *SyncExecuteTaskMonitorQueue) TimeSleepUntil() int64 {
	syncQueue.lock.Lock()
	defer syncQueue.lock.Unlock()
	if len(*(syncQueue.queue)) > 0 {
		return (*(syncQueue.queue))[0].when
	} else {
		return -1
	}
}

func (syncQueue *SyncExecuteTaskMonitorQueue) siftUpMonitor(i int) {

	if i >= len(*(syncQueue.queue)) {
		panic("queue data corruption")
	}
	when := (*(syncQueue.queue))[i].when
	tmp := (*(syncQueue.queue))[i]
	for i > 0 {

		p := (i - 1) / 4 // parent
		if when >= (*(syncQueue.queue))[p].when {
			break
		}
		(*(syncQueue.queue))[p].index = i
		(*(syncQueue.queue))[i] = (*(syncQueue.queue))[p]
		i = p
	}
	if tmp != (*(syncQueue.queue))[i] {
		tmp.index = i
		(*(syncQueue.queue))[i] = tmp
	}
}

func (syncQueue *SyncExecuteTaskMonitorQueue) siftDownMonitor(i int) {

	n := len(*(syncQueue.queue))
	if i >= n {
		panic("queue data corruption")
	}
	when := (*(syncQueue.queue))[i].when
	tmp := (*(syncQueue.queue))[i]
	for {
		c := i*4 + 1 // left child
		c3 := c + 2  // mid child
		if c >= n {
			break
		}
		w := (*(syncQueue.queue))[c].when
		if c+1 < n && (*(syncQueue.queue))[c+1].when < w {
			w = (*(syncQueue.queue))[c+1].when
			c++
		}
		if c3 < n {
			w3 := (*(syncQueue.queue))[c3].when
			if c3+1 < n && (*(syncQueue.queue))[c3+1].when < w3 {
				w3 = (*(syncQueue.queue))[c3+1].when
				c3++
			}
			if w3 < w {
				w = w3
				c = c3
			}
		}
		if w >= when {
			break
		}
		(*(syncQueue.queue))[c].index = i
		(*(syncQueue.queue))[i] = (*(syncQueue.queue))[c]
		i = c
	}
	if tmp != (*(syncQueue.queue))[i] {
		tmp.index = i
		(*(syncQueue.queue))[i] = tmp
	}
}
