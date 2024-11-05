package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i NotificationService -o ./mocks/ -s "_minimock.go"
