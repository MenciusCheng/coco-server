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
