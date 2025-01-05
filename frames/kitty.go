package frames

var Kitty = DefaultFrameType(kittyFrames)

var kittyFrames = []string{
`
⠀⠀⠀⠰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣁⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣄⡀⠀⢀⣀⣀⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣼⠝⠉⠉⠛⠷⣦⣤⣶⣶⣶⣶⡾⠃⠀⠈⢻⡾⠛⠉⠉⢻⡆⠀⠀⠀⠀
⠀⠀⠀⠀⣿⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⣼⠁⠀⣾⢻⡶⠛⠲⣄⣤⠼⢧⣄⠀⠀⠀
⠀⠀⠀⠀⣿⡾⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣄⠀⢙⣿⡀⠀⠀⣹⢶⡄⠀⣹⡀⠀⠀
⠀⠀⠀⣰⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠛⠒⢺⡛⠛⠁⣰⠏⠀⠀⠀
⠀⠀⢀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠒⠚⢹⡇⠀⠀⠀
⢠⡶⢼⣷⢤⠄⠀⠀⢀⣀⠀⠀⠀⠀⠀⠠⠀⠀⠀⠀⠀⣀⠀⠀⠀⠤⠾⣿⠖⠶⠆
⠀⠀⢈⣿⣀⡀⠀⠀⢿⡿⠀⠀⠀⠀⢀⣀⡀⠀⠀⠀⠸⣿⠇⠀⠀⠠⣴⣿⣄⣀⠀
⠀⠘⠋⠹⣷⣀⣀⠀⠀⠀⠀⠀⠀⠀⠯⣦⠽⠀⠀⠀⠀⠀⠀⠀⢀⣀⣾⠃⠀⠀⠀
⠀⠀⣠⡴⠟⠻⣧⡶⠳⣦⣠⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⡟⠋⠛⠶⠄⠀
⠀⠀⠁⠀⢀⡾⠛⠂⣀⡈⠀⢹⣧⣤⣄⣀⣀⣠⣤⣤⣴⣶⣾⠛⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠈⣷⣤⠄⠛⠃⠈⢿⢹⣄⠀⠀⠀⢠⡿⠀⠈⢷⡉⣻⣦⣤⣀⠀⠀⠀⠀
⠀⠀⠀⠀⣼⠏⢿⣤⣴⡷⣶⡟⠀⠙⠳⠶⠶⠛⠁⠀⠀⠈⣿⡋⠀⠈⣿⠆⠀⠀⠀
⠀⠀⠀⠀⢻⣆⠀⠀⢸⣿⣜⢷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣧⠀⠀⣽⠃⠀⠀⠀
⠀⠀⠀⠀⠀⠙⠻⣶⠞⠃⠙⠶⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⡶⠾⠋⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢸⡿⠶⢶⣤⣤⣤⣤⣤⣿⣤⣤⣤⣤⣤⣤⡶⢾⡇⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠸⣧⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⣸⡇⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠹⣷⣤⣀⣀⣀⣀⣠⣿⣀⣀⣀⣀⣀⣠⣴⠟⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠉⠉⠉⠛⠛⠛⠋⠉⠁⠀⠀⠀⠀⠀⠀`,
`
⣿⣿⣿⣏⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⠾⠻⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠛⠻⢿⣿⡿⠿⠿⢿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⠃⣢⣶⣶⣤⣈⠙⠛⠉⠉⠉⠉⢁⣼⣿⣷⡄⢁⣤⣶⣶⡄⢹⣿⣿⣿⣿
⣿⣿⣿⣿⠀⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⠃⣾⣿⠁⡄⢉⣤⣍⠻⠛⣃⡘⠻⣿⣿⣿
⣿⣿⣿⣿⠀⢁⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠻⣿⡦⠀⢿⣿⣿⠆⡉⢻⣿⠆⢿⣿⣿
⣿⣿⣿⠏⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣶⣶⣤⣭⡅⢤⣤⣾⠏⣰⣿⣿⣿
⣿⣿⡿⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣭⣥⡆⢸⣿⣿⣿
⡟⢉⡃⠈⡛⣻⣿⣿⡿⠿⣿⣿⣿⣿⣿⣟⣿⣿⣿⣿⣿⠿⣿⣿⣿⣛⣁⠀⣩⣉⣹
⣿⣿⡷⠀⠿⢿⣿⣿⡀⢀⣿⣿⣿⣿⡿⠿⢿⣿⣿⣿⣇⠀⣸⣿⣿⣟⠋⠀⠻⠿⣿
⣿⣧⣴⣆⠈⠿⠿⣿⣿⣿⣿⣿⣿⣿⣐⠙⣂⣿⣿⣿⣿⣿⣿⣿⡿⠿⠁⣼⣿⣿⣿
⣿⣿⠟⢋⣠⣄⠘⢉⣌⠙⠟⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠛⢠⣴⣤⣉⣻⣿
⣿⣿⣾⣿⡿⢁⣤⣽⠿⢷⣿⡆⠘⠛⠻⠿⠿⠟⠛⠛⠋⠉⠁⣤⣾⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣷⠈⠛⣻⣤⣼⣷⡀⡆⠻⣿⣿⣿⡟⢀⣿⣷⡈⢶⠄⠙⠛⠿⣿⣿⣿⣿
⣿⣿⣿⣿⠃⣰⡀⠛⠋⢈⠉⢠⣿⣦⣌⣉⣉⣤⣾⣿⣿⣷⠀⢴⣿⣷⠀⣹⣿⣿⣿
⣿⣿⣿⣿⡄⠹⣿⣿⡇⠀⠣⡈⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠘⣿⣿⠂⣼⣿⣿⣿
⣿⣿⣿⣿⣿⣦⣄⠉⣡⣼⣦⣉⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⢉⣁⣴⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⡇⢀⣉⡉⠛⠛⠛⠛⠛⠀⠛⠛⠛⠛⠛⠛⢉⡁⢸⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣇⠘⣿⣿⣿⣿⣿⣿⣿⠀⣿⣿⣿⣿⣿⣿⣿⠇⢸⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣆⠈⠛⠿⠿⠿⠿⠟⠀⠿⠿⠿⠿⠿⠟⠋⣠⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣶⣶⣶⣶⣶⣶⣤⣤⣤⣴⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿`,
}
