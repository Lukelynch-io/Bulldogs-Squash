#!/bin/zsh

session="bulldogs_site"

tmux new-session -d -s $session

tmux rename-window -t $session:1 "Site"
tmux send-keys -t "Site" 'nvim .' C-m
tmux new-window -t $session:2 -n 'bunrundev'
tmux send-keys -t 'bunrundev' 'bun run dev' C-m
