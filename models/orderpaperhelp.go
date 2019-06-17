package models

func translateProvinceID(name string) int {
	var IDs []int
	GetDb().Table("provinces").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translatePaperType(name string) int {
	var IDs []int
	GetDb().Table("paper_types").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateGrade(name string) int {
	var IDs []int
	GetDb().Table("grades").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateKeypoint(name string) int {
	var IDs []int
	GetDb().Table("keypoints").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateExerciseType(name string) int {
	var IDs []int
	GetDb().Table("exercise_types").Where("name = ?", name).Pluck("id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateIsOption(name string) int {
	switch name {
	case "单选题", "双选题", "多选题":
		return 1
	case "填空题":
		return 2
	case "计算题", "推断题", "简答题", "实验题":
		return 3
	default:
		return 0
	}
}
