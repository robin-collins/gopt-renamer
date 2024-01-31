[Setup]
AppName=gopt-renamer
AppVersion=0.1.0
DefaultDirName={pf}\gopt-renamer
OutputDir=..\
OutputBaseFilename=gopt-renamer-installer
Compression=lzma2
SolidCompression=yes
SetupIconFile="{#GetEnv('GITHUB_WORKSPACE')}\windows\icon.ico"
UninstallDisplayIcon="{app}\gopt-renamer.exe" 
AppPublisher="BlackCat-IT" 
VersionInfoCompany="BlackCat-IT"
VersionInfoDescription="use the power of AI to rename your images and screenshots."
VersionInfoVersion=0.1.0
DisableReadyMemo=True

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Files]
Source: "{#GetEnv('GITHUB_WORKSPACE')}\gopt-renamer.exe"; DestDir: "{app}"
Source: "{#GetEnv('GITHUB_WORKSPACE')}\gopt-renamer-contextmenu.exe"; DestDir: "{app}"

[Registry]
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer"; ValueType: string; ValueName: ""; ValueData: "Rename with gopt-renamer"; Flags: uninsdeletekey
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer\command"; ValueType: string; ValueName: ""; ValueData: """{app}\gopt-renamer-contextmenu.exe"" --image=""%1"" --force"; Flags: uninsdeletekey
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: """{app}\gopt-renamer-contextmenu.exe"",0"; Flags: uninsdeletekey

[Code]
#include "CustomPages.pas"