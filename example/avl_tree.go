package example

import (
	"awesomeProject5/base"
	"fmt"
)

// AvLTreeAdd
//
//	                                         19
//		                 ┌──────────────────────┴───────────────────────┐
//		                 7                                             69
//		     ┌───────────┴────────────┐                      ┌─────────┴──────────┐
//		     2                       11                     44                   95
//		┌────┴─────┐                 ┴─────┐           ┌────┴─────┐         ┌────┴─────┐
//		1          3                      14          21         57        85         99
//		           ┴──┐                                        ┌─┴──┐    ┌─┴──┐
//		              4                                       56   58   70   93
func AvLTreeAdd() {
	tree := base.NewBinaryTree(base.CmInt)
	arr := []int{85, 19, 69, 3, 7, 99, 95, 2, 1, 70, 44, 58, 11, 21, 14, 93, 57, 4, 56}
	for _, v := range arr {
		tree.Add(v)
	}
	tree.TreePrint()
}

func AvLTreeRemove() {
	tree := base.NewBinaryTree(base.CmInt)
	arr := []int{67, 52, 92, 96, 53, 95, 13, 63, 34, 82, 76, 54, 9, 68, 39}
	for _, v := range arr {
		tree.Add(v)
	}
	fmt.Println("=================================origin=================================")
	tree.TreePrint()
	for _, v := range arr {
		fmt.Printf("=================================remove:%v=================================\n", v)
		tree.Remove(v)
		tree.TreePrint()
	}
}

//=================================origin=================================
//                                                          67
//                                  ┌───────────────────────┴───────────────────────┐
//                                 52                                              82
//                     ┌───────────┴───────────┐                       ┌───────────┴───────────┐
//                    13                      54                      76                      95
//              ┌─────┴──────┐          ┌─────┴─────┐           ┌─────┴                 ┌─────┴─────┐
//              9           34         53          63          68                      92          96
//                          ┴──┐
//                            39
//
//=================================remove:67=================================
//                                                          68
//                                  ┌───────────────────────┴───────────────────────┐
//                                 52                                              82
//                     ┌───────────┴───────────┐                       ┌───────────┴───────────┐
//                    13                      54                      76                      95
//              ┌─────┴──────┐          ┌─────┴─────┐                                   ┌─────┴─────┐
//              9           34         53          63                                  92          96
//                          ┴──┐
//                            39
//
//=================================remove:52=================================
//                                                          68
//                                  ┌───────────────────────┴───────────────────────┐
//                                 53                                              82
//                     ┌───────────┴───────────┐                       ┌───────────┴───────────┐
//                    13                      54                      76                      95
//              ┌─────┴──────┐                ┴─────┐                                   ┌─────┴─────┐
//              9           34                     63                                  92          96
//                          ┴──┐
//                            39
//
//=================================remove:92=================================
//                                                          68
//                                  ┌───────────────────────┴───────────────────────┐
//                                 53                                              82
//                     ┌───────────┴───────────┐                       ┌───────────┴───────────┐
//                    13                      54                      76                      95
//              ┌─────┴──────┐                ┴─────┐                                         ┴─────┐
//              9           34                     63                                              96
//                          ┴──┐
//                            39
