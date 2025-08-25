run_dfs:
	go run assignment1_dfs.go assignment_helper.go
run_ids:
	go run assignment1_ids.go assignment_helper.go
run_main:
	go run assignment1_main.go assignment1_dfs.go assignment1_ids.go assignment1_biderctional.go assignment1_astar.go assignment_helper.go web_visualizer.go
run_main_UI:
	go run assignment1_main.go assignment1_dfs.go assignment1_ids.go assignment1_biderctional.go assignment1_astar.go assignment_helper.go web_visualizer.go --UI_ENABLED=true