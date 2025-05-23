package genstream

import (
	"encoding/json"
	"testing"
)

func Test_parseProtoService(t *testing.T) {
	protoContent := `
// 8周年主会场榜单
service ActivitySubAnniversary202505ExtObj {
  // 获取活动配置
  rpc FindActivityConfig(Anniversary202505_FindActivityConfigReq) returns (Anniversary202505_FindActivityConfigRes);
  // 查询组队榜单
  rpc FindRankManyPlayerSort(Anniversary202505_FindRankManyPlayerSortReq) returns (Anniversary202505_FindRankManyPlayerSortRes);
  // 查询家族榜
  rpc FindRankFamilySort(Anniversary202505_FindRankFamilySortReq) returns (Anniversary202505_FindRankFamilySortRes);
  // 查询家族榜房间贡献子榜
  rpc FindRankFamilySubRoomSort(Anniversary202505_FindRankFamilySubRoomSortReq) returns (Anniversary202505_FindRankFamilySubRoomSortRes);
  // 邀请cp
  rpc InviteCp(Anniversary202505_InviteCpReq) returns (Anniversary202505_InviteCpRes);
  // 邀请cp回复
  rpc InviteCpReply(Anniversary202505_InviteCpReplyReq) returns (Anniversary202505_InviteCpReplyRes);
  // 获取cp的信息回复
  rpc InviteCpPlayer(Anniversary202505_InviteCpPlayerReq) returns (Anniversary202505_InviteCpPlayerRes);
  // 获取好友列表cp组队情况
  rpc GetFriendCpInfo(Anniversary202505_FriendCpInfoReq) returns (Anniversary202505_FriendCpInfoRes);
  // 解绑cp
  rpc UnbindCp(Anniversary202505_UnbindCpReq) returns (Anniversary202505_UnbindCpRes);
  // 解绑cp回复
  rpc UnbindCpReply(Anniversary202505_UnbindCpReplyReq) returns (Anniversary202505_UnbindCpReplyRes);
}
`
	got := parseProtoService(protoContent)
	prettyJSON, _ := json.MarshalIndent(got, "", "  ")
	t.Logf("%s", string(prettyJSON))
}

func Test_parseProtoService2(t *testing.T) {
	protoContent := `
syntax = "proto3";
option objc_class_prefix = "PB3";

package act_pb;
import "act_pb/activity_define.proto";

// 新版奖励自定义奖励
message Genie202505_NewAward {
  int32 award_id = 1; // 奖励id
  string award_name = 2; // 奖励名称
  int32 prize_id = 3; // 奖品id
  string picture = 4; // 图片
  string remark = 5; // 备注
  int64 price = 6; // 单价
  int64 num = 7; // 数量
  string unit = 8; // 单位
  ActivityAwardType type = 9; // 奖励类型
  int32 effective_day = 10; // 有效时间（天）
  int32 send_type = 11;   // 时长发放类型 4当天过期、5子活动时间有效，其他值无效
}

// 主活动结构
message Genie202505_Activity {
  int32 id = 1; // 活动ID
  string name = 2; // 活动名称
  bool status = 3; // 活动状态，true为开启，false为关闭
  int64 start_time = 5; // 活动开始时间（时间戳）
  int64 end_time = 6; // 活动结束时间（时间戳）
  repeated Genie202505_RelActivity rel = 8; // 关联的子活动列表
  Genie202505_ConfActPage act_page = 9; // 页面内容配置
}

// 子活动
message Genie202505_RelActivity {
  int32 id = 1; // 子活动ID
  string name = 2; // 子活动名称
  int64 start_time = 6; // 子活动开始时间（时间戳）
  int64 end_time = 7; // 子活动结束时间（时间戳）
}

// 页面内容
message Genie202505_ConfActPage {
  int32 is_share = 1; // 支持分享0否1是
  map<string, string> act_rule = 2; // 活动规则kv
  string share_picture = 3; // 分享按钮图片
  map<string, Genie202505_ActPictureList> act_pic = 4; // 活动图片kv
}

message Genie202505_ActPictureList {
  repeated Genie202505_ActPicture pictures = 1;
}

message Genie202505_ActPicture {
  string name = 1;
  string picture = 2;
  string link = 3;
  int32 sort = 4; //排序
  int64 start_time = 5;
  int64 end_time = 6;
  int32 status = 7; //开启状态
  string video = 8;
}

// 玩家信息
message Genie202505_Player {
  int64 id = 1; // 用户ID
  string nickname = 2; // 昵称
  SexType sex = 3; // 性别
  int64 id2 = 4; // 靓号
  string icon = 5; // 图标
}

message Genie202505_BagItem {
  uint32 gift_id = 1; // 礼物id
  uint32 amount = 2; // 数量
}

message Genie202505_ConfLottery {
  int32 conf_id = 1; // 抽奖id
  string name = 2; // 名称
  int32 rel_act_id = 3; // 子活动
  int32 get_num_type = 4; // 获取抽奖次数类型 1-特定礼物 2-特定商品
  repeated Genie202505_AssignItem assign_items = 5; // 指定物品列表
  int32 lottery_point = 6; // 配置幸运值
  int32 refresh_price = 7; // 刷新价格
}

message Genie202505_AssignItem {
  int32 id = 1; // 是定物品id
  int32 num = 2; // 指定物品数量
}

message Genie202505_FindActivityConfigReq {
  int32 act_id = 1; // 活动id
}

message Genie202505_FindActivityConfigRes {
  Genie202505_Activity activity = 1; // 活动信息
  Genie202505_Player player_info = 2; // 查看个人信息
  repeated Genie202505_ConfLottery conf_lottery = 3; // 抽奖配置
  Genie202505_ActivityProcess activity_process = 4; // 活动进行信息
}

// 主活动进行信息
message Genie202505_ActivityProcess {
  int32 id = 1; // 活动ID
  ActivityProcessType process_status = 2; // 主活动进行状态，0未开始 1进行中 2已结束
  map<int32,Genie202505_RelActivityProcess> rel_process = 3; // 子活动id=>子活动状态
}

// 子活动进行信息
message Genie202505_RelActivityProcess {
  RelActivityProcessType rel_process_status = 1; // 子活动进行状态，0未开始 1进行中 2已结束
}

message Genie202505_FindLotteryUserStatusReq {
  int32 rel_act_id = 1; // 子活动ID
  int32 lottery_id = 2; // 抽奖ID
}

message Genie202505_FindLotteryUserStatusRes {
  repeated Genie202505_LotteryAward award_pool = 1; // 奖励池
  int32 lottery_lucky_point = 2; // 用户当前幸运值
  repeated Genie202505_BagItem items = 3; // 背包道具列表
  int64 gold = 4; // 金币
  repeated Genie202505_GenieBroadcast genie_broadcast = 5; // 本期幻化奖励轮播列表
}

message Genie202505_LotteryAward {
  int32 lottery_award_id = 1; // 奖池物品ID
  string lottery_award_name = 2; // 奖池物品名称
  int32 show_rate = 3; // 显示概率
  int32 weight = 4; // 奖品等级
  int64 start_time = 5;
  int64 end_time = 6;
  repeated Genie202505_NewAward awards = 7; // 奖励列表
  string tips = 8; // 奖品配置文案
  SexType sex = 9; // 性别
  int32 can_refresh = 10; // 是否可以刷新奖池，0否 1是
}

message Genie202505_GenieBroadcast {
  int64 conf_id = 1;
  string conf_name = 2;
  string conf_url = 3;
  int64 power = 4; // 灵力值
  int32 avatar_level = 5;  // 化身等级 1:初灵 2:半灵 3:地灵
  string avatar_level_icon = 6;  // 化身等级图标
}

// 抽奖
message Genie202505_DrawLuckyReq {
  int32 rel_act_id = 1; // 子活动id
  int32 num = 2; // 抽奖次数
  int32 lottery_id = 3; // 抽奖配置id
  SexType sex = 4; // 用户所选性别限制（该字段只有：随机奖池管理-物品配置-中选择性别限制才生效）
}

message Genie202505_DrawLuckyRes {
  repeated Genie202505_LuckyAward awards = 1; // 抽奖奖励列表
  int32 lottery_lucky_point = 2; // 用户当前幸运值
  repeated Genie202505_BagItem items = 3; // 背包道具列表
  int64 gold = 4; // 金币
}

// 抽奖奖励
message Genie202505_LuckyAward {
  int32 award_id = 1; // 奖励id
  string award_name = 2; // 奖励名称
  int32 prize_id = 3; // 奖品id
  string picture = 4; // 图片
  string remark = 5; // 备注
  int64 price = 6; // 单价
  int64 num = 7; // 数量
  string unit = 8; // 单位
  ActivityAwardType type = 9; // 奖励类型
  int32 effective_day = 10; // 有效天数，一般指除了特效以外的奖励的有效期
  int32 times = 11; // 抽奖时表示抽中的次数（单次奖励数量*times=num)
  int32 weight = 12; // 奖品等级
}

message Genie202505_FreshLotteryPoolReq {
  int32 rel_act_id = 1; // 子活动ID
  int32 lottery_id = 2; // 抽奖配置ID
  int32 lottery_award_id = 3; // 奖池物品ID
}

message Genie202505_FreshLotteryPoolRes {
  repeated Genie202505_LotteryAward award_pool = 1; // 奖励池
  int64 gold = 2; // 金币
}

// 抽奖记录
message Genie202505_ListLotteryRecordReq {
  int32 rel_act_id = 1; // 子活动id
  int32 conf_id = 2; // 抽奖配置id
  int32 page = 3; // 分页
  int32 page_size = 4; // 分页大小
}

message Genie202505_ListLotteryRecordRes {
  repeated Genie202505_LotteryRecord list = 1;
}

message Genie202505_LotteryRecord {
  int64 created_at = 1; // 获得时间
  string award_str = 2; // 奖励信息，奖励名称*数量+单位
}

message Genie202505_FindLotteryRefreshLogReq {
  int32 rel_act_id = 1; // 子活动id
  int32 page = 2; // 页码
  int32 page_size = 3; // 大小
}

message Genie202505_FindLotteryRefreshLogRes {
  repeated Genie202505_LotteryRefreshLog logs = 1;
}

message Genie202505_LotteryRefreshLog {
  int64 refresh_time = 1; // 刷新时间戳（秒）
  string award_name = 2; // 奖励名称
}

// 活动幻化202505期
service ActivitySubGenie202505ExtObj {
  // 获取活动配置
  rpc FindActivityConfig(Genie202505_FindActivityConfigReq) returns (Genie202505_FindActivityConfigRes);
  // 获取奖池用户状态
  rpc FindLotteryUserStatus(Genie202505_FindLotteryUserStatusReq) returns (Genie202505_FindLotteryUserStatusRes);
  // 抽奖
  rpc DrawLucky(Genie202505_DrawLuckyReq) returns (Genie202505_DrawLuckyRes);
  // 刷新奖池
  rpc FreshLotteryPool(Genie202505_FreshLotteryPoolReq) returns (Genie202505_FreshLotteryPoolRes);
  // 获取抽奖记录
  rpc ListLotteryRecord(Genie202505_ListLotteryRecordReq) returns (Genie202505_ListLotteryRecordRes);
  // 获取奖池刷新记录
  rpc FindLotteryRefreshLog(Genie202505_FindLotteryRefreshLogReq) returns (Genie202505_FindLotteryRefreshLogRes);
}
`
	got := parseProtoService(protoContent)
	prettyJSON, _ := json.MarshalIndent(got, "", "  ")
	t.Logf("%s", string(prettyJSON))
}
