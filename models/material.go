package models

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
)

func writeOMaterial() {
	abc()
	getMaterial("高中化学总复习版utf8.txt")
}

func getMaterial(name string) {
	paperBytes, _ := ioutil.ReadFile(KNOWLEDGE_BASE_URL + name)

	var test MaterialData
	err := json.Unmarshal(paperBytes, &test)
	if err != nil {
		beego.Debug(err)
		return
	}

	idCount := (int)(flakCurl.GetIntId())

	result := test.Result
	for r := range result {
		//录入数据
		var os OMaterialSimple
		os.ID = idCount
		os.Name = result[r].Name
		os.Parent = 0
		os.CreationDate = time.Now()
		os.ModificationDate = time.Now()
		GetDb().Table("materials").Create(&os)

		idCount = (int)(flakCurl.GetIntId())
		beego.Debug(idCount)
		beego.Debug(result[r].Parent)
		beego.Debug(len(result[r].Children))
		parent1ID := os.ID
		if result[r].Parent {
			firstResult := result[r].Children

			for f := range firstResult {
				var os OMaterialSimple
				os.ID = idCount
				os.Name = firstResult[f].Name
				os.Parent = parent1ID
				os.CreationDate = time.Now()
				os.ModificationDate = time.Now()
				GetDb().Table("materials").Create(&os)

				idCount = (int)(flakCurl.GetIntId())
				beego.Debug(idCount)
				parent2ID := os.ID
				if firstResult[f].Parent {
					secondResult := firstResult[f].Children

					for s := range secondResult {
						var os OMaterialSimple
						os.ID = idCount
						os.Name = secondResult[s].Name
						os.Parent = parent2ID
						os.CreationDate = time.Now()
						os.ModificationDate = time.Now()
						GetDb().Table("materials").Create(&os)

						idCount = (int)(flakCurl.GetIntId())
						beego.Debug(idCount)
						parent3ID := os.ID
						if secondResult[s].Parent {
							thirdResult := secondResult[s].Children

							for t := range thirdResult {
								var os OMaterialSimple
								os.ID = idCount
								os.Name = thirdResult[t].Name
								os.Parent = parent3ID
								os.CreationDate = time.Now()
								os.ModificationDate = time.Now()
								GetDb().Table("materials").Create(&os)

								idCount = (int)(flakCurl.GetIntId())
								beego.Debug(idCount)
								parent4ID := os.ID
								if thirdResult[t].Parent {
									fourResult := thirdResult[t].Children

									for a := range fourResult {
										var os OMaterialSimple
										os.ID = idCount
										os.Name = fourResult[a].Name
										os.Parent = parent4ID
										os.CreationDate = time.Now()
										os.ModificationDate = time.Now()
										GetDb().Table("materials").Create(&os)

										idCount = (int)(flakCurl.GetIntId())
										beego.Debug(idCount)
										parent5ID := os.ID
										if fourResult[a].Parent {
											fiveResult := fourResult[a].Children

											for b := range fiveResult {
												var os OMaterialSimple
												os.ID = idCount
												os.Name = fiveResult[b].Name
												os.Parent = parent5ID
												os.CreationDate = time.Now()
												os.ModificationDate = time.Now()
												GetDb().Table("materials").Create(&os)

												idCount = (int)(flakCurl.GetIntId())
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
