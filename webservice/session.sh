#!/bin/zsh

session="bulldogs_webservice"

tmux new-session -d -s $session

tmux rename-window -t 1 'Webservice'
tmux send-keys -t 'Webservice' 'nvim .' C-m
tmux new-window -t $session:2 -n 'CLI'
tmux send-keys -t 'CLI'
tmux attach -t $session
