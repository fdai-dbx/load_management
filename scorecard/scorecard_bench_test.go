package scorecard

import (
	"fmt"
	insecure_random "math/rand"
	"testing"
	"time"
)

func BenchmarkScorecard(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(16)
	rules := [][]Rule{
		{{"op:*;cat:*", 5}, {"cat:*", 5}, {"dog:*", 5}},
		{{"op:*;cat:*", 10}, {"dog:*", 5}},
		{{"cat:*", 5}, {"dog:*", 5}},
	}
	ops := []string{
		"select",
		"delete",
		"update",
		"insert",
	}
	sc := NewDynamicScorecard(rules[0])
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			op := ops[insecure_random.Intn(len(ops))]
			cat := insecure_random.Intn(5)
			dog := insecure_random.Intn(10)
			tags := []Tag{Tag("op:" + op), Tag(fmt.Sprintf("cat:%d", cat)), Tag(fmt.Sprintf("dog:%d", dog))}
			info := sc.TrackRequest(tags)
			if info.Tracked {
				time.Sleep(10 * time.Millisecond)
				info.Untrack()
			}
			// 10% of calls reconfigure scorecard
			if insecure_random.Float32() < 0.1 {
				newRules := rules[insecure_random.Intn(len(rules))]
				sc.Reconfigure(newRules)
			}
		}
	})
}

// BenchmarkScorecardGenerate exercises a realistic case where there are many
// tags matching a rule fragment and the production of AxB is costly.
func BenchmarkScorecardGenerate(b *testing.B) {
	b.ReportAllocs()
	sc := NewDynamicScorecard([]Rule{{"op:*;dog:*", 1}, {"op:*;cat:*", 5}, {"cat:*", 5}, {"dog:*", 5}})
	tags := []Tag{
		Tag("op:cat_create_txn"),
		Tag("dog:331"),
		Tag("dog:869"),
		Tag("dog:414"),
		Tag("dog:429"),
		Tag("dog:355"),
		Tag("dog:646"),
		Tag("dog:587"),
		Tag("dog:130"),
		Tag("dog:272"),
		Tag("dog:400"),
		Tag("dog:373"),
		Tag("dog:960"),
		Tag("dog:52"),
		Tag("dog:653"),
		Tag("dog:629"),
		Tag("dog:505"),
		Tag("dog:681"),
		Tag("dog:633"),
		Tag("dog:539"),
		Tag("dog:817"),
		Tag("dog:106"),
		Tag("dog:603"),
		Tag("dog:700"),
		Tag("dog:714"),
		Tag("dog:774"),
		Tag("dog:413"),
		Tag("dog:189"),
		Tag("dog:309"),
		Tag("dog:277"),
		Tag("dog:567"),
		Tag("dog:42"),
		Tag("dog:690"),
		Tag("dog:525"),
		Tag("dog:828"),
		Tag("dog:923"),
		Tag("dog:720"),
		Tag("dog:464"),
		Tag("dog:823"),
		Tag("dog:885"),
		Tag("dog:650"),
		Tag("dog:698"),
		Tag("dog:20"),
		Tag("dog:1010"),
		Tag("dog:593"),
		Tag("dog:131"),
		Tag("dog:121"),
		Tag("dog:318"),
		Tag("dog:805"),
		Tag("dog:22"),
		Tag("dog:972"),
		Tag("dog:920"),
		Tag("dog:640"),
		Tag("dog:147"),
		Tag("dog:228"),
		Tag("dog:585"),
		Tag("dog:531"),
		Tag("dog:668"),
		Tag("dog:645"),
		Tag("dog:7"),
		Tag("dog:265"),
		Tag("dog:227"),
		Tag("dog:808"),
		Tag("dog:334"),
		Tag("dog:493"),
		Tag("dog:910"),
		Tag("dog:744"),
		Tag("dog:1004"),
		Tag("dog:759"),
		Tag("dog:264"),
		Tag("dog:288"),
		Tag("dog:8"),
		Tag("dog:441"),
		Tag("dog:562"),
		Tag("dog:561"),
		Tag("dog:200"),
		Tag("dog:502"),
		Tag("dog:75"),
		Tag("dog:788"),
		Tag("dog:815"),
		Tag("dog:858"),
		Tag("dog:532"),
		Tag("dog:1002"),
		Tag("dog:524"),
		Tag("dog:507"),
		Tag("dog:812"),
		Tag("dog:15"),
		Tag("dog:23"),
		Tag("dog:122"),
		Tag("dog:197"),
		Tag("dog:939"),
		Tag("dog:799"),
		Tag("dog:763"),
		Tag("dog:19"),
		Tag("dog:248"),
		Tag("dog:879"),
		Tag("dog:634"),
		Tag("dog:726"),
		Tag("dog:82"),
		Tag("dog:659"),
		Tag("dog:90"),
		Tag("dog:735"),
		Tag("dog:105"),
		Tag("dog:930"),
		Tag("dog:155"),
		Tag("dog:848"),
		Tag("dog:13"),
		Tag("dog:182"),
		Tag("dog:91"),
		Tag("dog:110"),
		Tag("dog:56"),
		Tag("dog:733"),
		Tag("dog:886"),
		Tag("dog:420"),
		Tag("dog:321"),
		Tag("dog:632"),
		Tag("dog:527"),
		Tag("dog:435"),
		Tag("dog:31"),
		Tag("dog:622"),
		Tag("dog:226"),
		Tag("dog:138"),
		Tag("dog:465"),
		Tag("dog:137"),
		Tag("dog:553"),
		Tag("dog:96"),
		Tag("dog:756"),
		Tag("dog:993"),
		Tag("dog:710"),
		Tag("dog:74"),
		Tag("dog:942"),
		Tag("dog:826"),
		Tag("dog:268"),
		Tag("dog:477"),
		Tag("dog:898"),
		Tag("dog:702"),
		Tag("dog:1000"),
		Tag("dog:597"),
		Tag("dog:541"),
		Tag("dog:44"),
		Tag("dog:202"),
		Tag("dog:849"),
		Tag("dog:635"),
		Tag("dog:570"),
		Tag("dog:142"),
		Tag("dog:686"),
		Tag("dog:1003"),
		Tag("dog:651"),
		Tag("dog:386"),
		Tag("dog:784"),
		Tag("dog:1012"),
		Tag("dog:1005"),
		Tag("dog:417"),
		Tag("dog:933"),
		Tag("dog:461"),
		Tag("dog:749"),
		Tag("dog:589"),
		Tag("dog:2"),
		Tag("dog:900"),
		Tag("dog:956"),
		Tag("dog:902"),
		Tag("dog:119"),
		Tag("dog:352"),
		Tag("dog:855"),
		Tag("dog:1009"),
		Tag("dog:833"),
		Tag("dog:987"),
		Tag("dog:953"),
		Tag("dog:462"),
		Tag("dog:793"),
		Tag("dog:791"),
		Tag("dog:238"),
		Tag("dog:415"),
		Tag("dog:340"),
		Tag("dog:831"),
		Tag("dog:907"),
		Tag("dog:918"),
		Tag("dog:479"),
		Tag("dog:491"),
		Tag("dog:287"),
		Tag("dog:259"),
		Tag("dog:692"),
		Tag("dog:256"),
		Tag("dog:685"),
		Tag("dog:641"),
		Tag("dog:330"),
		Tag("dog:275"),
		Tag("dog:305"),
		Tag("dog:84"),
		Tag("dog:329"),
		Tag("dog:768"),
		Tag("dog:353"),
		Tag("dog:301"),
		Tag("dog:723"),
		Tag("dog:163"),
		Tag("dog:69"),
		Tag("dog:421"),
		Tag("dog:172"),
		Tag("dog:528"),
		Tag("dog:508"),
		Tag("dog:626"),
		Tag("dog:509"),
		Tag("dog:996"),
		Tag("dog:814"),
		Tag("dog:951"),
		Tag("dog:317"),
		Tag("dog:829"),
		Tag("dog:263"),
		Tag("dog:694"),
		Tag("dog:327"),
		Tag("dog:21"),
		Tag("dog:1006"),
		Tag("dog:969"),
		Tag("dog:187"),
		Tag("dog:706"),
		Tag("dog:26"),
		Tag("dog:1007"),
		Tag("dog:62"),
		Tag("dog:437"),
		Tag("dog:250"),
		Tag("dog:866"),
		Tag("dog:760"),
		Tag("dog:546"),
		Tag("dog:432"),
		Tag("dog:649"),
		Tag("dog:165"),
		Tag("dog:4"),
		Tag("dog:28"),
		Tag("dog:92"),
		Tag("dog:906"),
		Tag("dog:365"),
		Tag("dog:179"),
		Tag("dog:783"),
		Tag("dog:778"),
		Tag("dog:810"),
		Tag("dog:125"),
		Tag("dog:76"),
		Tag("dog:116"),
		Tag("dog:797"),
		Tag("dog:683"),
		Tag("dog:990"),
		Tag("dog:792"),
		Tag("dog:927"),
		Tag("dog:602"),
		Tag("dog:864"),
		Tag("dog:481"),
		Tag("dog:378"),
		Tag("dog:284"),
		Tag("dog:705"),
		Tag("dog:995"),
		Tag("dog:46"),
		Tag("dog:842")}
	for i := 0; i < b.N; i++ {
		sc.TrackRequest(tags)
	}
}

var benchmarkRules = []Rule{
	{"traffic:batch_traffic;source:blackbird_worker_prod_high_memory_sjc-blackbird_worker_bin", 5},
	{"source:segmentation*", 30},
	{"traffic:batch_traffic;source:filesystem.fs_job_worker_fs_job_worker-backfill_bin", 10},
	{"traffic:batch_traffic;tclass:master;client_id:*", 60},
	{"traffic:batch_traffic;tclass:slave;source:owner_team_data_infra-mapper.py", 5},
	{"traffic:batch_traffic;tclass:slave;client_id:*", 5},
	{"traffic:live_traffic;source:metaserver*", 400},
	{"traffic:live_traffic;source:*", 50},
	{"op:gid_create_txn", 100},
	{"op:gid_create_txn;colo:*", 1},
	{"traffic:batch_traffic;tclass:master;source:*", 30},
	{"traffic:batch_traffic;tclass:slave;source:*", 60},
	{"op:read_list;source:cape*", 20},
	{"op:scan", 10},
	{"source:*;op:scan", 2}}
var requests = [][]Tag{
	{"source:cape_yss_workers_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda_iad-async_task_worker_wrapper.py", "client_id:meta", "op:insert_revision", "op:read_gid2", "traffic:live_traffic"},
	{"source:fsverifier_fsverifier_worker_prod-fsverifier_worker_bin", "client_id:audit_log", "op:read_revision", "op:txn", "traffic:live_traffic"},
	{"source:cape_dispatcher_yss_edgestore_canary-cape_dispatcher_bin", "client_id:cape_dispatcher", "op:gid_create_read_id", "op:read_ep", "traffic:batch_traffic"},
	{"source:owner_shared_spaces-count_shmodels_banned_under_rl_migration.py", "client_id:argus", "op:read_xtxn_conflict", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:filejournal_sjc.prod-fj_server_1.12_bin", "client_id:taskrunner_sharing_platform", "op:txn", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:blackbird_prod_common", "op:txn", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:owner_messaging_team-delphi_store_consumer_bin", "client_id:megaphone_bluemail_kafka_consumer", "op:read_gid2", "op:insert_revision", "traffic:batch_traffic"},
	{"source:metaserver_courier_live_site_control-main.py", "client_id:megaphone_bluemail_kafka_consumer", "op:insert_revision", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:sync_frontend_sjc.canary-sync_frontend_server_bin", "client_id:atf_hsc", "op:gid_create_read_id", "op:read_gid2", "traffic:live_traffic"},
	{"source:audit_log.atf_async_file_events_logging_workers_download_file_event_lambda-atf_controller_bin", "client_id:blackbird_prod_common", "op:read_revision", "op:insert_revision", "traffic:live_traffic"},
	{"source:metaserver_client-paster", "client_id:cape_dispatcher", "op:read_xtxn_conflict", "op:insert_revision", "traffic:live_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:filesystem", "op:gid_create_read_id", "op:insert_revision", "traffic:batch_traffic"},
	{"source:owner_shared_spaces-count_shmodels_banned_under_rl_migration.py", "client_id:taskrunner_team_lifecycle", "op:read_xtxn_conflict", "op:read_revision", "traffic:batch_traffic"},
	{"source:metaserver_courier_live_site_control-main.py", "client_id:atf_controller", "op:gid_create_read_id", "op:insert_revision", "traffic:batch_traffic"},
	{"source:atf.frontend_yss_frontend_prod-atf_frontend_bin", "client_id:taskrunner_sharing_platform", "op:gid_create_read_id", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:audit_log.atf_async_file_events_logging_workers_download_file_event_lambda-atf_controller_bin", "client_id:blackbird_prod_alki-qa-streamer", "op:insert_revision", "op:read_revision", "traffic:live_traffic"},
	{"source:fs_move_worker_sjc.prod-fs_move_worker_bin", "client_id:argus", "op:read_ep", "op:read_xtxn_conflict", "traffic:live_traffic"},
	{"source:filejournal_sjc.prod-fj_server_1.12_bin", "client_id:atf_hsc", "op:gid_create_read_id", "op:read_gid2", "traffic:live_traffic"},
	{"source:metaserver_courier_live_site-main.py", "client_id:segmentation", "op:read_gid2", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:owner_filesystem_team-namespace_member_backedge_consistency_checker.py", "client_id:sync_frontend", "op:read_gid2", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:filejournal_sjc.prod-fj_server_1.12_bin", "client_id:megaphone_journey_builder", "op:read_xtxn_conflict", "op:prepare_xtxn", "traffic:live_traffic"},
	{"source:megaphone_megaphone_rpc_service_prod-megaphone_rpc_service.py", "client_id:megaphone_journey_builder", "op:insert_revision", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:metaserver_courier_live_site-main.py", "client_id:team_lifecycle", "op:read_ep", "op:insert_revision", "traffic:batch_traffic"},
	{"source:metaserver_client_control-paster", "client_id:cloud_docs", "op:read_xtxn_conflict", "op:insert_revision", "traffic:batch_traffic"},
	{"source:fsverifier_fsverifier_worker_prod-fsverifier_worker_bin", "client_id:iam", "op:prepare_xtxn", "op:read_gid2", "traffic:live_traffic"},
	{"source:audit_log.atf_async_file_events_logging_workers_download_file_event_lambda-atf_controller_bin", "client_id:megaphone_bluemail_kafka_consumer", "op:read_xtxn_conflict", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:atf_controller", "op:gid_create_read_id", "op:txn", "traffic:live_traffic"},
	{"source:fsverifier_fsverifier_worker_prod-fsverifier_worker_bin", "client_id:filesystem", "op:read_revision", "op:read_revision", "traffic:batch_traffic"},
	{"source:cape_yss_workers_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda_iad-async_task_worker_wrapper.py", "client_id:cape_workers", "op:gid_create_read_id", "op:read_gid2", "traffic:batch_traffic"},
	{"source:owner_messaging_team-delphi_store_consumer_bin", "client_id:audit_log", "op:read_revision", "op:read_revision", "traffic:batch_traffic"},
	{"source:atf.store_consumer_yss_store_consumer_prod-atf_store_consumer_bin", "client_id:sprinkle", "op:gid_create_read_id", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:atf.store_consumer_yss_store_consumer_prod-atf_store_consumer_bin", "client_id:taskrunner_sharing_platform", "op:gid_create_read_id", "op:gid_create_read_id", "traffic:live_traffic"},
	{"source:audit_log.atf_async_file_events_logging_workers_download_file_event_lambda-atf_controller_bin", "client_id:meta", "op:txn", "op:insert_revision", "traffic:batch_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:argus", "op:insert_revision", "op:txn", "traffic:batch_traffic"},
	{"source:atf.store_consumer_yss_store_consumer_prod-atf_store_consumer_bin", "client_id:atf_store_consumer", "op:insert_revision", "op:read_gid2", "traffic:batch_traffic"},
	{"source:cape_yss_workers_PhotoAndVideoCollectionsRefreshTopology_PhotoVideoCollectionsRefreshLambda-pv_collections_refresh.py", "client_id:argus", "op:read_revision", "op:insert_revision", "traffic:live_traffic"},
	{"source:audit_log.atf_async_file_events_logging_workers_download_file_event_lambda-atf_controller_bin", "client_id:dropbox", "op:txn", "op:insert_revision", "traffic:live_traffic"},
	{"source:blackbird_worker_prod_restorations_sjc-blackbird_worker_bin", "client_id:filesystem", "op:read_xtxn_conflict", "op:txn", "traffic:batch_traffic"},
	{"source:sync_frontend_sjc.prod-sync_frontend_server_bin", "client_id:filesystem", "op:gid_create_read_id", "op:txn", "traffic:batch_traffic"},
	{"source:metaserver_client-paster", "client_id:segmentation", "op:insert_revision", "op:read_ep", "traffic:batch_traffic"},
	{"source:filejournal_sjc.prod-fj_server_1.12_bin", "client_id:iam", "op:gid_create_read_id", "op:insert_revision", "traffic:live_traffic"},
	{"source:sync_frontend_sjc.prod-sync_frontend_server_bin", "client_id:taskrunner_team_lifecycle", "op:insert_revision", "op:txn", "traffic:batch_traffic"},
	{"source:owner_filesystem_team-namespace_member_backedge_consistency_checker.py", "client_id:blackbird_prod_common", "op:insert_revision", "op:read_xtxn_conflict", "traffic:live_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:megaphone_journey_builder", "op:read_xtxn_conflict", "op:read_gid2", "traffic:batch_traffic"},
	{"source:megaphone_megaphone_rpc_service_prod-megaphone_rpc_service.py", "client_id:cloud_docs", "op:read_revision", "op:gid_create_read_id", "traffic:live_traffic"},
	{"source:cape_yss_workers_PhotoAndVideoCollectionsRefreshTopology_PhotoVideoCollectionsRefreshLambda-pv_collections_refresh.py", "client_id:blackbird_prod_restorations", "op:gid_create_read_id", "op:read_xtxn_conflict", "traffic:live_traffic"},
	{"source:owner_messaging_team-delphi_store_consumer_bin", "client_id:audit_log", "op:insert_revision", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:megaphone_megaphone_rpc_service_prod-megaphone_rpc_service.py", "client_id:taskrunner_team_lifecycle", "op:read_xtxn_conflict", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:owner_shared_spaces-count_shmodels_banned_under_rl_migration.py", "client_id:taskrunner_prod", "op:read_xtxn_conflict", "op:read_xtxn_conflict", "traffic:live_traffic"},
	{"source:sync_frontend_sjc.prod-sync_frontend_server_bin", "client_id:atf_hsc", "op:read_xtxn_conflict", "op:insert_revision", "traffic:batch_traffic"},
	{"source:filejournal_sjc.prod-fj_server_1.12_bin", "client_id:argus", "op:prepare_xtxn", "op:read_revision", "traffic:live_traffic"},
	{"source:metaserver_client_control-paster", "client_id:atf_hsc", "op:read_gid2", "op:read_gid2", "traffic:live_traffic"},
	{"source:gatekeeper_sjc.prod-gatekeeper_svc_bin", "client_id:blackbird_prod_high-memory", "op:read_xtxn_conflict", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:cape_dispatcher_yss_edgestore_prod-cape_dispatcher_bin", "client_id:blackbird_prod_alki-qa-streamer", "op:txn", "op:read_gid2", "traffic:batch_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:megaphone_bluemail_kafka_consumer", "op:read_xtxn_conflict", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:blackbird_prod_common", "op:prepare_xtxn", "op:read_gid2", "traffic:batch_traffic"},
	{"source:cape_yss_workers_PhotoAndVideoCollectionsRefreshTopology_PhotoVideoCollectionsRefreshLambda-pv_collections_refresh.py", "client_id:taskrunner_prod", "op:gid_create_read_id", "op:insert_revision", "traffic:live_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:dropbox", "op:read_ep", "op:insert_revision", "traffic:live_traffic"},
	{"source:atf.store_consumer_yss_store_consumer_prod-atf_store_consumer_bin", "client_id:atf_controller", "op:prepare_xtxn", "op:read_gid2", "traffic:batch_traffic"},
	{"source:cape_dispatcher_yss_edgestore_canary-cape_dispatcher_bin", "client_id:dropbox", "op:txn", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:metaserver_courier_live_site-main.py", "client_id:atf_hsc", "op:gid_create_read_id", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:atf_test_cluster_workers_atf_test_lambda-atf_controller_bin", "client_id:argus", "op:gid_create_read_id", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:gatekeeper_sjc.prod-gatekeeper_svc_bin", "client_id:megaphone_bluemail_kafka_consumer", "op:read_revision", "op:read_gid2", "traffic:batch_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:search_indexer", "op:read_gid2", "op:gid_create_read_id", "traffic:live_traffic"},
	{"source:sync_frontend_sjc.prod-sync_frontend_server_bin", "client_id:taskrunner_team_lifecycle", "op:prepare_xtxn", "op:read_xtxn_conflict", "traffic:live_traffic"},
	{"source:gatekeeper_sjc.prod-gatekeeper_svc_bin", "client_id:atf_store_consumer", "op:read_ep", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:owner_messaging_team-delphi_store_consumer_bin", "client_id:taskrunner_team_lifecycle", "op:gid_create_read_id", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:atf_test_cluster_workers_atf_test_lambda-atf_controller_bin", "client_id:megaphone_journey_builder", "op:txn", "op:txn", "traffic:live_traffic"},
	{"source:fs_move_worker_sjc.prod-fs_move_worker_bin", "client_id:blackbird_prod_high-memory", "op:read_xtxn_conflict", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:fs_move_worker_sjc.prod-fs_move_worker_bin", "client_id:cloud_docs", "op:gid_create_read_id", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:filesystem.fs_job_worker_fs_job_worker-backfill_bin", "client_id:taskrunner_prod", "op:read_ep", "op:read_revision", "traffic:batch_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:blackbird_prod_high-memory", "op:prepare_xtxn", "op:read_revision", "traffic:live_traffic"},
	{"source:fsverifier_fsverifier_worker_prod-fsverifier_worker_bin", "client_id:sprinkle", "op:read_ep", "op:read_revision", "traffic:live_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:filesystem", "op:txn", "op:txn", "traffic:live_traffic"},
	{"source:metaserver_client_control-paster", "client_id:argus", "op:gid_create_read_id", "op:read_revision", "traffic:live_traffic"},
	{"source:fs_move_worker_sjc.prod-fs_move_worker_bin", "client_id:sprinkle", "op:gid_create_read_id", "op:prepare_xtxn", "traffic:live_traffic"},
	{"source:atf.store_consumer_yss_store_consumer_prod-atf_store_consumer_bin", "client_id:dropbox", "op:prepare_xtxn", "op:read_revision", "traffic:batch_traffic"},
	{"source:metaserver_client-paster", "client_id:argus", "op:prepare_xtxn", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:metaserver_courier_live_site_control-main.py", "client_id:search_indexer", "op:read_revision", "op:insert_revision", "traffic:batch_traffic"},
	{"source:cape_yss_workers_PhotoAndVideoCollectionsRefreshTopology_PhotoVideoCollectionsRefreshLambda-pv_collections_refresh.py", "client_id:argus", "op:read_xtxn_conflict", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:owner_filesystem_team-namespace_member_backedge_consistency_checker.py", "client_id:iam", "op:prepare_xtxn", "op:read_ep", "traffic:batch_traffic"},
	{"source:blackbird_worker_prod_restorations_sjc-blackbird_worker_bin", "client_id:cloud_docs", "op:read_gid2", "op:read_ep", "traffic:batch_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:blackbird_prod_high-memory", "op:read_revision", "op:read_revision", "traffic:batch_traffic"},
	{"source:filesystem.fs_job_worker_fs_job_worker-backfill_bin", "client_id:taskrunner_team_lifecycle", "op:read_xtxn_conflict", "op:read_ep", "traffic:batch_traffic"},
	{"source:metaserver_courier_live_site_control-main.py", "client_id:taskrunner_prod", "op:read_xtxn_conflict", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:megaphone_megaphone_rpc_service_prod-megaphone_rpc_service.py", "client_id:dropbox", "op:prepare_xtxn", "op:gid_create_read_id", "traffic:live_traffic"},
	{"source:metaserver_courier_live_site_control-main.py", "client_id:atf_store_consumer", "op:prepare_xtxn", "op:read_gid2", "traffic:batch_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:cape_dispatcher", "op:read_xtxn_conflict", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:cape_yss_workers_canary_asyncTaskWorkerWrapperTopology_asyncTaskWorkerWrapperLambda-async_task_worker_wrapper.py", "client_id:atf_controller", "op:prepare_xtxn", "op:prepare_xtxn", "traffic:batch_traffic"},
	{"source:megaphone_megaphone_rpc_service_prod-megaphone_rpc_service.py", "client_id:team_lifecycle", "op:read_gid2", "op:prepare_xtxn", "traffic:live_traffic"},
	{"source:metaserver_client-paster", "client_id:atf_store_consumer", "op:read_xtxn_conflict", "op:read_ep", "traffic:batch_traffic"},
	{"source:owner_shared_spaces-count_shmodels_banned_under_rl_migration.py", "client_id:taskrunner_sharing_platform", "op:read_ep", "op:read_revision", "traffic:live_traffic"},
	{"source:sync_frontend_sjc.canary-sync_frontend_server_bin", "client_id:atf_store_consumer", "op:read_gid2", "op:prepare_xtxn", "traffic:live_traffic"},
	{"source:owner_messaging_team-delphi_store_consumer_bin", "client_id:blackbird_prod_alki-qa-streamer", "op:insert_revision", "op:read_xtxn_conflict", "traffic:batch_traffic"},
	{"source:metaserver_client-paster", "client_id:audit_log", "op:insert_revision", "op:read_revision", "traffic:batch_traffic"},
	{"source:owner_filesystem_team-namespace_member_backedge_consistency_checker.py", "client_id:taskrunner_prod", "op:read_xtxn_conflict", "op:gid_create_read_id", "traffic:batch_traffic"},
	{"source:metaserver_client_canary-paster", "client_id:cloud_docs", "op:gid_create_read_id", "op:txn", "traffic:live_traffic"},
}

func BenchmarkBucket(b *testing.B) {
	b.ReportAllocs()
	sc := NewDynamicScorecard([]Rule{{"op:*;dog:*", 1}, {"op:*;cat:*", 5}, {"cat:*", 5}, {"dog:*", 5}})
	tag := Tag("op:cat_create_txn")
	for i := 0; i < b.N; i++ {
		_ = sc.(*scorecardImpl).bucket(tag)
	}
}

func BenchmarkProdDataSetWithRelease(b *testing.B) {
	b.ReportAllocs()
	benchmarkSC := NewScorecard(benchmarkRules)
	for i := 0; i < b.N; i++ {
		for _, r := range requests {
			trackingInfo := benchmarkSC.TrackRequest(r)
			go func(ti *TrackingInfo) {
				time.Sleep(1 * time.Millisecond)
				ti.Untrack()
			}(trackingInfo)
		}
	}
}

func BenchmarkProdDataSetWithoutRelease(b *testing.B) {
	b.ReportAllocs()
	benchmarkSC := NewScorecard(benchmarkRules)
	for i := 0; i < b.N; i++ {
		for _, r := range requests {
			benchmarkSC.TrackRequest(r)
		}
	}
}
