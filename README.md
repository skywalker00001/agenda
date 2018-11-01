# Agenda

`entity`负责创建实体与操作实体，`cmd`负责程序的逻辑

## entity

### User

1. `NewUser`创建一个`User`实体并返回它的指针 

```go
NewUser func(name, password, email, phone string) *User
```

2. `xxx.Getxxx()`及`xxx.Setxxx(xxx)`用来获取或设置`User`实体的属性

```go
// Getxxx
GetName func() string
GetPassword func() string
GetEmail func() string
GetPhone func() string
// Setxxx
SetName func(newName string)
SetPassword func(newPassword string)
SetEmail func(newEmail string)
SetPhone func(newPhone string)
```

### Date

1. `NewDate`创建一个`Date`实体并返回它的指针

```go
NewDate func(year, month, day, hour, minute int) *Date
```

2. `xxx.Getxxx()`及`xxx.Setxxx(xxx)`**方法**用来获取或设置`Date`实体的属性

```go
// Getxxx
GetYear func() int
GetMonth func() int
GetDay func() int
GetHour func() int
GetMinute func() int
// Setxxx
SetYear func(newYear int) 
SetMonth func(newMonth int)
SetDay func(newDay int) 
SetHour func(newHour int) 
SetMinute func(newMinute int) 
```

3. `xxx.IsValid()`判断日期是否合法

```go
IsValid func() bool
```

4. `StringToDate()`及`DateToString()`**函数**用来执行`string`和`Date`类型相互转化

```go
//StringToDate convert a date string to a Date type
StringToDate func(dateString string) Date
// DateToString convert a Date struct to a string with format YYYY-MM-DD/HH:mm
DateToString func(date Date) string
```

5. 因为没有操作符重载，所以设计`xxx.IsEqual(xxx)`, `xxx.IsGreater(xxx)`, `xxx.IsLess(xxx)`, `xxx.IsGreaterThanEqual(xxx)`及`xxx.IsLessThanEqual(xxx)`**方法**来判断日期的前后关系

```go
IsEqual func(date Date) bool
IsGreater func(date Date) bool
IsLess func(date Date) bool
IsGreaterThanEqual func(date Date) bool
IsLessThanEqual func(date Date) bool
```

### Meeting

1. `NewMeeting`创建一个`Meeting`实体并返回它的指针

```go
NewMeeting func(sponsor string, title string, startTime Date, endTime Date, participators []string) *Meeting
```

2. `xxx.Getxxx()`及`xxx.Setxxx(xxx)`**方法**用来获取或设置`Meeting`实体的属性

```go
// Getxxx
GetSponsor func() string
GetTitle func() string
GetStartTime func() Date
GetEndTime func() Date
GetParticipators func() []string
// Setxxx
SetSponsor func(sponsor string)
SetTitle func(title string)
SetStartTime func(startTime Date)
SetEndTime func(endTime Date)
SetParticipators func(participators []string)
```

3. `xxx.AddParticipator(xxx)`为`meeting`实体添加指定参与者xxx

```go
AddParticipator func(participator string)
```

4. `xxx.RemoveParticipator（xxx)`为`meeting`实体删除指定参与者xxx

```go
RemoveParticipator func(participator string)
```

5. `xxx.IsParticipator(xxx)`判断xxx是否在`meeting`实体中

```go
IsParticipator func(username string) bool
```

### Storage

**注意：Storage的方法如果调用成功会自动写入`curUser.txt`, `Meeting.json`和`User.json`**

1. `GetStorage()`**函数**获取一个单例`storage`实体读写`curUser.txt`, `Meeting.json`和`User.json`

```go
GetStorage func() *Storage
```

2. `xxx.CreateUser(xxx)`**方法**用来创建新用户

```go
CreateUser func(newUser User)
```

3. `xxx.QueryUser(uFilter)`**方法**通过传入一个过滤器`uFilter`来**查询用户**

```go
// uFilter
type uFilter func(*User) bool
// QueryUser *
QueryUser func(filter uFilter) []User
```

4. `xxx.UpdateUser(uFilter, uSwitcher)`**方法**通过传入一个过滤器`uFilter`筛选用户并对它们使用`uSwitcher`进行**更新**

```go
type uFilter func(*User) bool
type uSwitcher func(*User)
// UpdateUser *
UpdateUser func(filter uFilter, switcher uSwitcher) int
```

5. `xxx.DeleteUser(uFilter)`**方法**通过传入一个过滤器`uFilter`筛选用户并**删除**

```go
// uFilter
type uFilter func(*User) bool
// DeleteUser *
DeleteUser func(filter uFilter) int
```

6. `xxx.CreateMeeting(xxx)`**方法**用来创建新会议并写入`Meeting.json`

```go
CreateMeeting func(newMeeting Meeting)
```

7. `xxx.QueryMeeting(mFilter)`**方法**通过传入一个过滤器`mFilter`来**查询会议**

```go
// mFilter
type mFilter func(*Meeting) bool
// QueryMeeting *
QueryMeeting func(filter mFilter) []Meeting
```

8. `xxx.UpdateMeeting(mFilter, mSwitcher)`**方法**通过传入一个过滤器`mFilter`筛选会议并对它们使用`mSwitcher`进行**更新**

```go
type mFilter func(*Meeting) bool
type mSwitcher func(*Meeting)
// UpdateMeeting *
UpdateMeeting func(filter mFilter, switcher mSwitcher) int
```

9. `xxx.DeleteMeeting(mFilter)`方法通过传入一个过滤器`mFilter`筛选会议并**删除**

```go
// mFilter
type mFilter func(*Meeting) bool
// DeleteUser *
DeleteMeeting func(filter mFilter) int
```

10. `xxx.GetCurUser()`获取现在登录的用户

```go
SetCurUser func(u User)
```

11. `xxx.SetCurUser(xxx)`设置当前登录用户

```go
SetCurUser func(u User)
```

### Logger

`NewLogger`根据传入的前缀来创建新的`logger`，其可同时写入`log.txt`和显示在屏幕上

```go
// NewLogger create a logger which write info on screen and in ./data/log.txt with specific prefix
NewLogger func(prefix string) *log.Logger
```

