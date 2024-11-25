#!/bin/zsh

session="bulldogs"

tmux new-session -d -s $session

tmux rename-window -t 1 'Webservice'
tmux send-keys -t 'Webservice' 'cd webservice' C-m 'clear' C-m
tmux new-window -t $session:2 -n 'CLI'
tmux new-window -t $session:3 -n 'Site'
tmux send-keys -t 'Site' 'cd site' C-m 'clear' C-m
tmux new-window -t $session:4 -n 'bunrundev'
tmux send-keys -t 'bunrundev' 'cd site && bun run dev' C-m
tmux attach -t bulldogs
