package models

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego"
)

const (
	KNOWLEDGE_BASE_URL = "./知识点文件/"
)

func writeKnowledge() {
	abc()
	// rd, _ := ioutil.ReadDir("./知识点文件/")

	// for r := range rd {
	// 	getKnowledge(rd[r].Name())
	// }

	getKnowledge("初中英语知识点.txt")
	// writePaperSql("2017-2018学年河南省驻马店市确山县七年级（下）期末数学试卷.txt")

}

func getKnowledge(name string) {
	paperBytes, _ := ioutil.ReadFile(KNOWLEDGE_BASE_URL + name)

	var test KnowledgeData
	err := json.Unmarshal(paperBytes, &test)
	if err != nil {
		beego.Debug(err)
		return
	}

	idCount := 1 + 2200000

	result := test.Result
	for r := range result {
		//录入数据
		var keyPoint KeyPoint
		keyPoint.KeyPointID = idCount
		keyPoint.Name = result[r].Name
		GetDb().Table("t_keypoints").Create(&keyPoint)

		idCount++
		beego.Debug(idCount)
		beego.Debug(result[r].Parent)
		beego.Debug(len(result[r].Children))

		if result[r].Parent {
			firstResult := result[r].Children

			for f := range firstResult {
				var keyPoint KeyPoint
				keyPoint.KeyPointID = idCount
				keyPoint.Name = firstResult[f].Name
				GetDb().Table("t_keypoints").Create(&keyPoint)

				idCount++
				beego.Debug(idCount)

				if firstResult[f].Parent {
					secondResult := firstResult[f].Children

					for s := range secondResult {
						var keyPoint KeyPoint
						keyPoint.KeyPointID = idCount
						keyPoint.Name = secondResult[s].Name
						GetDb().Table("t_keypoints").Create(&keyPoint)

						idCount++
						beego.Debug(idCount)

						if secondResult[s].Parent {
							thirdResult := secondResult[s].Children

							for t := range thirdResult {
								var keyPoint KeyPoint
								keyPoint.KeyPointID = idCount
								keyPoint.Name = thirdResult[t].Name
								GetDb().Table("t_keypoints").Create(&keyPoint)

								idCount++
								beego.Debug(idCount)

								if thirdResult[t].Parent {
									fourResult := thirdResult[t].Children

									for a := range fourResult {
										var keyPoint KeyPoint
										keyPoint.KeyPointID = idCount
										keyPoint.Name = fourResult[a].Name
										GetDb().Table("t_keypoints").Create(&keyPoint)

										idCount++
										beego.Debug(idCount)

										if fourResult[a].Parent {
											fiveResult := fourResult[a].Children

											for b := range fiveResult {
												var keyPoint KeyPoint
												keyPoint.KeyPointID = idCount
												keyPoint.Name = fiveResult[b].Name
												GetDb().Table("t_keypoints").Create(&keyPoint)

												idCount++
												beego.Debug(idCount)

												if fiveResult[b].Parent {
													beego.Debug("还有你的妹妹")
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
