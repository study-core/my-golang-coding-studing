package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"math/big"
)

func main() {


	subDelegateFn := func(source, sub *big.Int) (*big.Int, *big.Int) {

		return new(big.Int).Sub(source, sub), common.Big0
	}

	refundFn := func(remain, aboutRelease, aboutLockRepo *big.Int) (*big.Int, *big.Int, *big.Int, bool, error) {
		// When remain is greater than or equal to del.ReleasedTmp/del.Released
		if remain.Cmp(common.Big0) > 0 {
			if remain.Cmp(aboutRelease) >= 0 && aboutRelease.Cmp(common.Big0) > 0 {

				remain, aboutRelease = subDelegateFn(remain, aboutRelease)

			} else if remain.Cmp(aboutRelease) < 0 {
				// When remain is less than or equal to del.ReleasedTmp/del.Released
				aboutRelease, remain = subDelegateFn(aboutRelease, remain)
			}
		}

		if remain.Cmp(common.Big0) > 0 {

			// When remain is greater than or equal to del.LockRepoTmp/del.LockRepo
			if remain.Cmp(aboutLockRepo) >= 0 && aboutLockRepo.Cmp(common.Big0) > 0 {


					return remain, aboutRelease, aboutLockRepo, false, nil
				}



			} else if remain.Cmp(aboutLockRepo) < 0 {
				// When remain is less than or equal to del.LockRepoTmp/del.LockRepo


			return remain, aboutRelease, aboutLockRepo, false, nil


		}

		return remain, aboutRelease, aboutLockRepo, true, nil
	}


	del := &staking.Delegation{
		ReleasedHes:        common.Big2,
		RestrictingPlanHes: common.Big1,
	}

	//remain, release, lockRepo := common.Big2, common.Big1, common.Big3

	remain := common.Big3

	fmt.Println(remain, del.ReleasedHes, del.RestrictingPlanHes)

	// 注意 结构体的不能直接赋值？
	remain, release, lock, flag, err := refundFn(remain, del.ReleasedHes, del.RestrictingPlanHes)

	del.ReleasedHes, del.RestrictingPlanHes = release, lock

	fmt.Println(remain, del.ReleasedHes, del.RestrictingPlanHes, flag, err)

}
