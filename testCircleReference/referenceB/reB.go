package referenceB

import "myCryto-study/testCircleReference/referenceA"

func BTest1(){
	referenceA.Atest2()
}

func BTest2(){
	//referenceA.FuncCh <- BTest1
	//close(referenceA.FuncCh)
	referenceA.Func = BTest1
}
