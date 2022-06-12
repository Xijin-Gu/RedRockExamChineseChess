package service

import (
	"ChineseChess/dao"
	"ChineseChess/model"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func JudgeGameUser(userName string,gameID int)(bool,model.Game){
	db := dao.Link()
	defer db.Close()
	gameData := dao.QueryGameDataByID(db,gameID)
	if (userName == gameData.Mover || userName == gameData.Waiter) {
		return true,gameData
	}
	return false,gameData
}



func ParseCheckboard(str string)[10][9]string{
	var checkboards [10][18]string
	var checkboard [10][9]string
	a := strings.Split(str,"\n")
	for i,v := range a {
		for i1,v1 := range v {
			checkboards[i][i1] = string(v1)
		}
	}
	for i:=0;i<10;i++{
		for j:=0;j<9;j++{
			checkboard[i][j] = checkboards[i][2*j]+checkboards[i][2*j+1]
		}
	}
	return checkboard
}

func MoveChess(start,end string,checkboard [10][9]string,gameData model.Game)(model.Game){
	startCode,_ := strconv.Atoi(start)
	endCode,_ := strconv.Atoi(end)
	si := startCode/10-1
	sj := startCode%10-1
	ei := endCode/10-1
	ej := endCode%10-1
	//确定棋子类型
	fmt.Println(startCode)
	fmt.Println(startCode%10)
	fmt.Println(si," ",sj," ",ei," ",ej)
	pieceT := checkboard[si][sj]
	if pieceT == "00" {
		return gameData
	}
	//验证走棋逻辑
	var boM  bool
	switch pieceT {
	case "11":
		boM = KingMove(startCode,endCode,"11")
		break
	case "21":
		boM = KingMove(startCode,endCode,"21")
		break
	case "12":
		boM = GuardMove(startCode,endCode,"12")
		break
	case "22":
		boM = GuardMove(startCode,endCode,"22",)
		break
	case "13":
		boM = BishopMove(startCode,endCode,"13",checkboard)
		break
	case "23":
		boM = BishopMove(startCode,endCode,"23",checkboard)
		break
	case "14":
		boM = KnightMove(startCode,endCode,checkboard)
		break
	case "24":
		boM = KnightMove(startCode,endCode,checkboard)
		break
	case "15":
		boM = RookAndCannonMove(startCode,endCode,checkboard)
		break
	case "25":
		boM = RookAndCannonMove(startCode,endCode,checkboard)
		break
	case "16":
		boM = RookAndCannonMove(startCode,endCode,checkboard)
		break
	case "26":
		boM = RookAndCannonMove(startCode,endCode,checkboard)
		break
	case "17":
		boM = PawnMove(startCode,endCode,"17")
		break
	case "27":
		boM = PawnMove(startCode,endCode,"27")
		break
	}
	if !boM {
		return gameData
	}
	//验证是否已有棋子
	piecejs,_ := strconv.Atoi(pieceT)
	pieceje,_ := strconv.Atoi(checkboard[ei][ej])
	if piecejs%10 == pieceje%10 {
		return gameData
	}
	//验证吃子
	if pieceT == "16" || pieceT == "26" {
		boM = CannonEat(startCode,endCode,checkboard)
		if !boM {
			return gameData
		}
	}
	checkboard[ei][ej] = checkboard[si][sj]
	checkboard[si][sj] = "00"
	var checkboardStr string
	for i:=0;i<10;i++{
		for j:=0;j<9;j++{
			checkboardStr = checkboardStr+checkboard[i][j]
		}
		checkboardStr = checkboardStr + "\n"
	}
	gameData.Checkerboard = checkboardStr
	return gameData

}
func KingMove(start,end int,code string)bool{
	si := start/10-1
	sj := start%10-1
	ei := end/10-1
	ej := end%10-1
	//- 满足3<j<7,
	//- 编码为11，满足1<=i<=3
	//- 编码为21，满足8<=i<=10
	//- 满足|si+sj-ei-ej|==1
	if !(sj<7 && sj>3 && 3<ej && ej<7){
		return false
	}
	if code == "11" {
		if !(si >= 1 && si <=3 && ei >= 1 && ei <= 3) {
			return false
		}
	} else {
		if !(si >= 8 && si <=10 && ei >= 8 && ei <= 10) {
			return false
		}
	}
	if !(math.Abs(float64(si+sj-ei-ej)) == 1) {
		return false
	}
	return true


}
func GuardMove(start,end int,code string)bool{
	si := start/10-1
	sj := start%10-1
	ei := end/10-1
	ej := end%10-1
	//- 满足3<i<7,
	//	- 编码为12，满足1<=i<=3
	//- 编码为22，满足8<=i<=10
	//- 满足|si*10+sj-ei*10+ej|==11
	if !(sj<7 && sj>3 && 3<ej && ej<7){
		return false
	}
	if code == "12" {
		if !(si >= 1 && si <=3 && ei >= 1 && ei <= 3) {
			return false
		}
	} else {
		if !(si >= 8 && si <=10 && ei >= 8 && ei <= 10) {
			return false
		}
	}
	if !(math.Abs(float64(start-end)) == 11) {
		return false
	}
	return true

}

func BishopMove(start,end int,code string,checkboard [10][9]string)bool{
	si := start/10-1
	sj := start%10-1
	ei := end/10-1
	ej := end%10-1
	//- 编码为13，满足1<=i<=5
	//- 编码为23，满足6<=i<=10
	//- 满足|si*10+sj-ei*10-ej|==22
	//- 且(start+end)/2处值为00
	if code == "13" {
		if !(si >= 1 && si <=5 && ei >= 1 && ei <= 5) {
			return false
		}
	} else {
		if !(si >= 6 && si <=10 && ei >= 6 && ei <= 10) {
			return false
		}
	}
	if !(math.Abs(float64(start-end)) == 22) {
		return false
	}
	if checkboard[(si+ei)/2][(sj+ej)/2] != "00" {
		return false
	}
	return true


}

func KnightMove(start,end int,checkboard [10][9]string)bool{
	si := start/10-1
	sj := start%10-1
	//- |end-start|== 12 || 21
	//- 马腿为，|end-start|== 12,(si,sj+1)不能由棋子；|end-start|== 21，（si+1,sj)处不能由棋子。
	if !(math.Abs(float64(start-end)) == 12 || math.Abs(float64(start-end)) == 21) {
		return false
	}
	if math.Abs(float64(start-end)) == 12 {
		if checkboard[si][sj+1] != "00" {
			return false
		}
	} else {
		if checkboard[si+1][sj] != "00" {
			return false
		}
	}
	return true

}

func RookAndCannonMove(start,end int,checkboard [10][9]string)bool{
	si := start/10-1
	sj := start%10-1
	ei := end/10-1
	ej := end%10-1
	//- i,j只有一个能动。
	//- 开始结束之间不能有棋子
	if (si != ei && sj != ej){
		return false
	}
	if (si == ei){
		for j:=sj+1;j<ej;j++{
			if checkboard[ei][j] != "00"{
				return false
			}
		}
	} else {
		for i:=si+1;i<ej;i++{
			if checkboard[i][sj] != "00"{
				return false
			}
		}
	}
	return true

}



func PawnMove(start,end int,code string)bool{
	si := start/10-1
	sj := start%10-1
	ei := end/10-1
	ej := end%10-1
	//- ei>=si
	//
	//- 编码为17，i<=5,sj=ej
	//- 编码为27，i>=6,sj=ej
	//- |end-start| == 1 || 10
	if ei < si {
		return false
	}
	if code == "17" {
		if (si <= 5){
			if (sj != ej){
				return false
			}
		}
	} else {
		if (si >= 6){
			if (sj != ej){
				return false
			}
		}
	}
	if !(math.Abs(float64(start-end)) == 1) {
		return false
	}
	return true
}

func CannonEat(start,end int,checkboard [10][9]string)bool{
	si := start/10-1
	sj := start%10-1
	ei := end/10-1
	ej := end%10-1
	//- ei,ej处编码和si,sj处不同
	//- 遍历中间的值，存在且仅存在一个非00值
	pieceTs,_ := strconv.Atoi(checkboard[si][sj])
	pieceTe,_ := strconv.Atoi(checkboard[ei][ej])
	pieceTs = pieceTs%10
	pieceTe = pieceTe%10
	if pieceTs == pieceTe || pieceTe == 0 {
		return false
	}
	var v int
	if si == ei {
		for j:=sj+1;j<ej;j++{
			if checkboard[si][j] == "00" {
				v++
			}
		}
	}
	if v != 1 {
		return false
	}
	return true

}



func UpdateChessData(game model.Game)string{
	db := dao.Link()

	dao.UpdateGameData(db,game)
	gameData := dao.QueryGameDataByID(db,game.ID)
	defer db.Close()
	return gameData.Checkerboard
}


