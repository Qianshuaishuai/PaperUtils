package models

func translateProvinceIDForReadboy(name string) int {
	var IDs []int
	GetDb().Table("t_provinces").Where("F_name = ?", name).Pluck("F_province_id", &IDs)

	if len(IDs) > 0 {
		return IDs[0]
	}

	return 0
}

func translateGradeForReadboy(name string) int {
	switch name {
	case "高一":
		return 10
	case "高二":
		return 11
	case "高三":
		return 12
	case "七年级":
		return 7
	case "八年级":
		return 8
	case "九年级":
		return 9
	case "一年级":
		return 1
	case "二年级":
		return 2
	case "三年级":
		return 3
	case "四年级":
		return 4
	case "五年级":
		return 5
	case "六年级":
		return 6

	default:
		return 0
	}
}

func translateSemesterForReadboy(name string) int {
	switch name {
	case "高一":
		return 4
	case "高二":
		return 2
	case "高三":
		return 1
	case "七年级":
		return 12
	case "八年级":
		return 10
	case "九年级":
		return 8
	case "一年级":
		return 24
	case "二年级":
		return 22
	case "三年级":
		return 20
	case "四年级":
		return 18
	case "五年级":
		return 16
	case "六年级":
		return 14

	default:
		return 0
	}
}
