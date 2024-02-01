:: Add context menu entry for gopt-renamer
reg add "HKCR\SystemFileAssociations\image\shell\gopt-renamer" /ve /t REG_SZ /d "Rename with gopt-renamer" /f
:: Add icon
reg add "HKCR\SystemFileAssociations\image\shell\gopt-renamer" /v "Icon" /t REG_SZ /d "\"{app}\\gopt-renamer-contextmenu.exe\",0" /f
:: Add command
reg add "HKCR\SystemFileAssociations\image\shell\gopt-renamer\command" /ve /t REG_SZ /d "\"D:\\ai\\gopt-renamer\\gopt-renamer.exe\" --image=\"%1\" --silent" /f
:: Add icon for default
reg add "HKCR\SystemFileAssociations\image\shell\gopt-renamer\DefaultIcon" /ve /t REG_SZ /d "\"{app}\\gopt-renamer-contextmenu.exe\",0" /f
:: Change Windows Explorer Default Behavior for Multiple Files Selected to change from a maximum of 15 to 32
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer" /v MultipleInvokePromptMinimum /t REG_DWORD /d 32 /f
