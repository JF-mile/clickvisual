package factory

//
// func Test_dayTime2Timestamp(t *testing.T) {
// 	type args struct {
// 		in string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int64
// 	}{
// 		{
// 			name: "test-1",
// 			args: args{
// 				in: "2022-01-11T17:39:49+08:00",
// 			},
// 			want: 1641893989,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := factory.dayTime2Timestamp(tt.args.in, "'2006-01-02T15:04:05+08:00'"); got != tt.want {
// 				t.Errorf("dayTime2Timestamp() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_queryTransformer(t *testing.T) {
// 	type args struct {
// 		in string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantOut string
// 		wantErr bool
// 	}{
// 		{
// 			name: "test-1",
// 			args: args{
// 				in: "_namespace_='kube-system' and _log_agent_='fluent-bit-8w7qh' and _time_second_='2022-01-11T17:39:49+08:00'",
// 			},
// 			wantOut: "_namespace_='kube-system' and _log_agent_='fluent-bit-8w7qh' and _time_second_='1641893989'",
// 			wantErr: false,
// 		}, {
// 			name: "test-2",
// 			args: args{
// 				in: "_namespace_='kube-system'",
// 			},
// 			wantOut: "_namespace_='kube-system'",
// 			wantErr: false,
// 		}, {
// 			name: "test-3",
// 			args: args{
// 				in: "_namespace_ like '%kube-system%'",
// 			},
// 			wantOut: "_namespace_ like '%kube-system%'",
// 			wantErr: false,
// 		}, {
// 			name: "test-4",
// 			args: args{
// 				in: "_namespace_ = '=====kube-system%'",
// 			},
// 			wantOut: "_namespace_ = '=====kube-system%'",
// 			wantErr: false,
// 		}, {
// 			name: "test-5",
// 			args: args{
// 				in: "reqAid = 'androidxlv'",
// 			},
// 			wantOut: "reqAid = 'androidxlv'",
// 			wantErr: false,
// 		}, {
// 			name: "test-6",
// 			args: args{
// 				in: "andreqAid = 'androidxlv'",
// 			},
// 			wantOut: "andreqAid = 'androidxlv'",
// 			wantErr: false,
// 		}, {
// 			name: "test-7",
// 			args: args{
// 				in: "==_namespace_ = '=====kube-system%'",
// 			},
// 			wantOut: "==_namespace_ = '=====kube-system%'",
// 			wantErr: false,
// 		}, {
// 			name: "test-8",
// 			args: args{
// 				in: "==_namespace_ = '=====kube-system%' and andreqAid = 'xx and roidxlv'",
// 			},
// 			wantOut: "==_namespace_ = '=====kube-system%' and andreqAid = 'xx and roidxlv'",
// 			wantErr: false,
// 		}, {
// 			name: "test-9",
// 			args: args{
// 				in: "1='1'",
// 			},
// 			wantOut: "1='1'",
// 			wantErr: false,
// 		}, {
// 			name: "test-10",
// 			args: args{
// 				in: "1=1",
// 			},
// 			wantOut: "1=1",
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotOut, err := clickhouse.queryTransformer(tt.args.in)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("queryTransformer() got %v, error %v, wantErr %v", gotOut, err, tt.wantErr)
// 				return
// 			}
// 			if gotOut != tt.wantOut {
// 				t.Errorf("queryTransformer() gotOut = %v, want %v", gotOut, tt.wantOut)
// 			}
// 		})
// 	}
// }
