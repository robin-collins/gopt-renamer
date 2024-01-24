reg add "HKCR\SystemFileAssociations\image\shell\gopt-renamer" /ve /t REG_SZ /d "Rename with gopt-renamer" /f
reg add "HKCR\SystemFileAssociations\image\shell\gopt-renamer\command" /ve /t REG_SZ /d "\"D:\\ai\\gopt-renamer\\gopt-renamer.exe\" --image=\"%1\" --silent" /f
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer" /v MultipleInvokePromptMinimum /t REG_DWORD /d 32 /f
