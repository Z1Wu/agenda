package entity

// 处理和会议有关的逻辑

import (
	"fmt"
)

// 创建会议
/*
已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册 用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
	会议主题(title)（在会议列表中具有唯一性）
	会议参与者(participator)
	会议起始时间(startDate)
	会议结束时间(endDate)
1.注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）
  与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
2.用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。
*/
func CreateMeeting(userName, title, startDate, endDate string, participator []string) bool {
	var start Date
	var end Date
	start = buildDateFromString(startDate)
	end = buildDateFromString(endDate)

	if (!start.isValid()) || (!end.isValid()) {
		fmt.Println("日期不合法")
		return false
	}
	if start.GreaterOrEqual(end) {
		fmt.Println("开始日期不可大于或等于结束日期")
		return false
	}

	/*-----------检查发起人是否登陆用户----------------*/
	if userName != CurrentUser.getName() {
		fmt.Println("发起人非登陆用户")
		return false
	}

	/*----------------是否包含未注册用户--------------*/
	userfilter := func(u *User) bool {
		for _, p := range participator {
			if p == u.getName() {
				return true
			}
		}
		return false
	}
	userlist := queryUser(userfilter)
	if len(userlist) != len(participator) {
		fmt.Println("存在参与者未注册")
		return false
	}

	/*-------------检查会议主题是否存在----------------*/
	meetingfilter1 := func(m *Meeting) bool {
		return title == m.getTitle()
	}
	meetinglist1 := queryMeeting(meetingfilter1)
	if len(meetinglist1) != 0 {
		fmt.Println("会议主题已存在")
		return false
	}

	/*-------------是否与发起人其他会议冲突----------------*/
	meetingfilter2 := func(m *Meeting) bool {
		if !(userName == m.getSponsor() || m.isParticipator(userName)) {
			return false
		}
		if (userName == m.getSponsor() || m.isParticipator(userName)) &&
			(start.GreaterOrEqual(buildDateFromString(m.getEndDate())) ||
				end.SmallerOrEqual(buildDateFromString(m.getStartDate()))) {
			return false
		} else {
			return true
		}
	}
	meetinglist2 := queryMeeting(meetingfilter2)
	if len(meetinglist2) != 0 {
		fmt.Println("与发起人其他会议冲突")
		return false
	}

	/*---------------是否与参与者其他会议冲突---------------*/
	meetingfilter3 := func(m *Meeting) bool {
		for _, p := range participator {
			if !(p == m.getSponsor() || m.isParticipator(p)) {
				return false
			}
			if (p == m.getSponsor() || m.isParticipator(p)) &&
				(start.GreaterOrEqual(buildDateFromString(m.getEndDate())) ||
					end.SmallerOrEqual(buildDateFromString(m.getStartDate()))) {
				return false
			} else {
				return true
			}
		}
		return false
	}
	meetinglist3 := queryMeeting(meetingfilter3)
	if len(meetinglist3) != 0 {
		fmt.Println("与参与者其他会议冲突")
		return false
	}

	/*------------参与者是否重复-------------*/

	for i := 0; i < len(participator); i++ {
		for j := i + 1; j < len(participator); j++ {
			if participator[i] == participator[j] {
				fmt.Println("参与者不能重复")
				return false
			}
		}
	}

	/*--------------参与者中是否有发起者--------------------*/
	for _, p := range participator {
		if userName == p {
			fmt.Println("参与者中不能有发起者")
			return false
		}
	}

	/*-----------参与者是否为空-----------------*/
	if len(participator) == 0 {
		fmt.Println("参与者不能为空")
		return false
	}

	createMeeting(Meeting{userName, title, startDate, endDate, participator})
	return true
}

/*
*查询会议
*
*已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
*用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
*在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
*注意，查询会议的结果应包括用户作为 发起者或参与者 的会议
 */
func MeetingQuery(sponsor, startDate, endDate string) []Meeting {
	var temp []Meeting
	start := buildDateFromString(startDate)
	end := buildDateFromString(endDate)
	if start.isMoreThan(end) || !start.isValid() || !end.isValid() {
		fmt.Println("日期不合法")
		return temp //此时a为空
	}

	filter := func(a *Meeting) bool {
		if (a.Sponsor == sponsor || a.isParticipator(sponsor)) &&
			(buildDateFromString(a.getEndDate()).GreaterOrEqual(start) && buildDateFromString(a.getStartDate()).SmallerOrEqual(start)) {
			return true
		}
		if (a.Sponsor == sponsor || a.isParticipator(sponsor)) &&
			(buildDateFromString(a.getStartDate()).SmallerOrEqual(end)) && buildDateFromString(a.getStartDate()).GreaterOrEqual(start) {
			return true
		}
		return false
	}
	return queryMeeting(filter)
}

/* list all meetings the user sponsor */
func ListAllSponsorMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return name == a.Sponsor
	}
	return queryMeeting(filter)
}

/*
取消会议
1.已登录的用户可以取消 自己发起 的某一会议安排。
2.取消会议时，需提供唯一标识：会议主题（title）。
*/
func DeleteMeeting(name, title string) bool {
	if name != CurrentUser.getName() {
		fmt.Println("非登陆用户")
		return false
	}
	filter := func(a *Meeting) bool {
		return a.Title == title && a.Sponsor == name
	}
	return deleteMeeting(filter) > 0
}

/*
清空会议
1.已登录的用户可以清空 自己发起 的所有会议安排。
*/
func DeleteAllMeetings(name string) bool {
	if name != CurrentUser.getName() {
		fmt.Println("非登陆用户")
		return false
	}
	if len(ListAllSponsorMeetings(name)) == 0 {
		fmt.Println("该用户没有发起的会议")
		return false
	}
	filter := func(a *Meeting) bool {
		return a.Sponsor == name
	}
	return deleteMeeting(filter) > 0
}

/*
增删会议参与者
1.已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
2.增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
3.删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。
*/
/* 增加参会者 */
func Addparticipator(title string, participator []string) bool {

	if len(participator) == 0 {
		fmt.Println("必须至少添加一个参与者")
		return false
	}

	filter1 := func(m *Meeting) bool {
		return m.Sponsor == CurrentUser.Name && m.Title == title
	}
	//找到创建人的一条会议
	mlist := queryMeeting(filter1)
	if len(mlist) == 0 {
		fmt.Println("找不到当前用户创建的该会议")
		return false
	}

	for i := 0; i < len(participator); i++ {
		for j := i + 1; j < len(participator); j++ {
			if participator[i] == participator[j] {
				fmt.Println("添加的参与者不能重复")
				return false
			}
		}
	}

	for _, p := range mlist[0].Participators {
		for _, pp := range participator {
			if pp == p {
				fmt.Println("存在添加的参与者已参与该会议")
				return false
			}
		}
	}

	for _, p := range participator {
		if CurrentUser.Name == p {
			fmt.Println("参与者不能有发起者")
			return false
		}
	}

	start := buildDateFromString(mlist[0].StartDate)
	end := buildDateFromString(mlist[0].EndDate)
	filter2 := func(m *Meeting) bool {
		for _, p := range participator {
			if !(p == m.getSponsor() || m.isParticipator(p)) {
				return false
			}
			if (p == m.getSponsor() || m.isParticipator(p)) &&
				(start.GreaterOrEqual(buildDateFromString(m.getEndDate())) ||
					end.SmallerOrEqual(buildDateFromString(m.getStartDate()))) {
				return false
			} else {
				return true
			}
		}
		return true
	}
	mlist1 := queryMeeting(filter2)
	if len(mlist1) != 0 {
		fmt.Println("添加的参与者与其他会议冲突")
		return false
	}

	//add
	for _, p := range participator {
		mlist[0].Participators = append(mlist[0].Participators, p)
	}

	//写回原会议
	for i, m := range meetingList {
		if m.Title == mlist[0].Title {
			meetingList[i] = mlist[0]
		}
	}

	return true
}

/* 删除参会者 */
func Removeparticipator(title string, participator []string) bool {

	if len(participator) == 0 {
		fmt.Println("必须至少删除一个参与者")
		return false
	}

	filter1 := func(m *Meeting) bool {
		return m.Sponsor == CurrentUser.Name && m.Title == title
	}
	//找到创建人的一条会议
	mlist := queryMeeting(filter1)
	if len(mlist) == 0 {
		fmt.Println("找不到当前用户创建的该会议")
		return false
	}

	for _, p := range participator {
		if CurrentUser.Name == p {
			fmt.Println("不能删除发起者,你可以退出会议（quit）")
			return false
		}
	}

	if len(participator) > len(mlist[0].Participators) {
		fmt.Println("删除人数不能大于原有人数")
		return false
	}

	for _, p := range participator {
		if mlist[0].isParticipator(p) == false {
			fmt.Println("存在待删除者没有参与该会议")
			return false
		}
	}

	//delete participator
	for _, p := range participator {
		n := 0
		for i, pp := range mlist[0].Participators {
			if p == pp {
				mlist[0].Participators[i] = mlist[0].Participators[len(mlist[0].Participators)-1-n]
				n++
			}
		}
		mlist[0].Participators = mlist[0].Participators[:len(mlist[0].Participators)-n]
	}

	//重新写回meetingList
	for i, m := range meetingList {
		if m.Title == mlist[0].Title {
			meetingList[i] = mlist[0]
		}
	}

	filter2 := func(m *Meeting) bool {
		if len(m.Participators) == 0 {
			return true
		}
		return false
	}
	if deleteMeeting(filter2) == 1 {
		fmt.Println("该会议的参与者已清空，会议已删除")
		return true
	}

	return true
}

/*
退出会议
1.已登录的用户可以退出 自己参与 的某一会议安排。
2.退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。
*/
func QuitMeeting(title string) bool {

	filter1 := func(m *Meeting) bool {
		return (m.Sponsor == CurrentUser.Name || m.isParticipator(CurrentUser.Name)) && m.Title == title
	}
	//找到当前用户参加的一条会议，当前用户为参与者或发起者
	mlist := queryMeeting(filter1)
	if len(mlist) == 0 {
		fmt.Println("找不到当前用户参加的该会议")
		return false
	}

	//当前用户为该会议的发起者，则删除该会议
	if CurrentUser.Name == mlist[0].Sponsor {
		if DeleteMeeting(CurrentUser.Name, title) {
			fmt.Println("当前用户为该会议的发起者，该会议已被删除")
			return true
		}
	}

	//delete CurrentUser
	n := 0
	for i, pp := range mlist[0].Participators {
		if CurrentUser.Name == pp {
			mlist[0].Participators[i] = mlist[0].Participators[len(mlist[0].Participators)-1-n]
			n++
		}
	}
	mlist[0].Participators = mlist[0].Participators[:len(mlist[0].Participators)-n]

	//重新写回meetingList
	for i, m := range meetingList {
		if m.Title == mlist[0].Title {
			meetingList[i] = mlist[0]
		}
	}

	filter2 := func(m *Meeting) bool {
		if len(m.Participators) == 0 {
			return true
		}
		return false
	}
	if deleteMeeting(filter2) == 1 {
		fmt.Println("该会议的参与者已清空，会议已删除")
		return true
	}

	return true
}
