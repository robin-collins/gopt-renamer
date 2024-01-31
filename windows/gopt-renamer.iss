[Setup]
AppName=gopt-renamer
AppVersion=1.0.0
DefaultDirName={pf}\gopt-renamer
OutputDir=..\
OutputBaseFilename=gopt-renamer-installer

[Files]
Source: "{#GetEnv('GITHUB_WORKSPACE')}\gopt-renamer.exe"; DestDir: "{app}"

[Registry]
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer"; ValueType: string; ValueName: ""; ValueData: "Rename with gopt-renamer"
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer\command"; ValueType: string; ValueName: ""; ValueData: """{app}\gopt-renamer.exe"" --image=""%1"" --force"
