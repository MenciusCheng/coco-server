package generator

import (
	"reflect"
	"testing"
)

func TestCalSqlHead(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *SqlHead
	}{
		{
			args: args{
				line: "CREATE TABLE `miza_activity`.`act_auction_gift_conf` (",
			},
			want: &SqlHead{
				TableName: "act_auction_gift_conf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalSqlHead(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalSqlHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalSqlField(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *SqlField
	}{
		{
			args: args{
				line: "`id` int(11) AUTO_INCREMENT COMMENT 'ID:自增主键',",
			},
			want: &SqlField{
				ColName: "id",
				ColType: "int(11)",
				Comment: "ID:自增主键",
			},
		},
		{
			args: args{
				line: "`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			},
			want: &SqlField{
				ColName: "updated_at",
				ColType: "datetime",
				Comment: "更新时间",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalSqlField(tt.args.line, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalSqlField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalSqlFoot(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *SqlFoot
	}{
		{
			args: args{
				line: ") COMMENT='活动礼物竞拍配置';",
			},
			want: &SqlFoot{
				Comment: "活动礼物竞拍配置",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalSqlFoot(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalSqlFoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParserSQL2(t *testing.T) {
	text := "CREATE TABLE `miza_activity`.`act_auction_gift_conf` (\n    `id` int(11) AUTO_INCREMENT COMMENT 'ID:自增主键',\n    `config_name` varchar(255) NOT NULL DEFAULT '' COMMENT '配置名',\n    `act_id` int(11) NOT NULL DEFAULT '0' COMMENT '主活动id',\n    `rel_act_id` int(11) NOT NULL DEFAULT '0' COMMENT '子活动id',\n    `start_time` int(11) NOT NULL DEFAULT '0' COMMENT '开始时间',\n    `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束时间',\n    `extend` json COMMENT '扩展配置',\n    `operator` varchar(128) NOT NULL DEFAULT '' COMMENT '操作人',\n    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n    PRIMARY KEY (`id`)\n) COMMENT='活动礼物竞拍配置';"
	got := ParserSQL2(text, nil)
	t.Logf("%+v", got)
}
