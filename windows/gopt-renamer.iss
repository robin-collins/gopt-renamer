[Setup]
AppName=gopt-renamer
AppVersion=0.1.0
DefaultDirName={pf}\gopt-renamer
OutputDir=..\
OutputBaseFilename=gopt-renamer-installer
Compression=lzma2
SolidCompression=yes
SetupIconFile="icon.ico" ; Ensure the icon file is present in the script directory
VersionInfoCompany="BlackCat-IT"
VersionInfoDescription="use the power of AI to rename your images and screenshots."
VersionInfoVersion=0.1.0
DisableReadyMemo=True

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Files]
Source: "{#GetEnv('GITHUB_WORKSPACE')}\gopt-renamer.exe"; DestDir: "{app}"

[Registry]
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer"; ValueType: string; ValueName: ""; ValueData: "Rename with gopt-renamer"
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer\command"; ValueType: string; ValueName: ""; ValueData: """{app}\gopt-renamer.exe"" --image=""%1"" --force"

[Code]
var
  OpenAIKeyPage: TWizardPage;
  OpenAIKeyEdit: TEdit;

function OpenAIKeyValid(Key: string): Boolean;
begin
  // Check if the key is not empty and starts with 'sk-'
  Result := (Length(Key) > 0) and (Copy(Key, 1, 3) = 'sk-');
end;

procedure InitializeWizard();
begin
  OpenAIKeyPage := CreateCustomPage(wpWelcome, 'OpenAI API Key', 'Please enter your OpenAI API Key (it should start with ''sk-'')');
  OpenAIKeyEdit := TEdit.Create(WizardForm);
  OpenAIKeyEdit.Parent := OpenAIKeyPage.Surface;
  OpenAIKeyEdit.Top := ScaleY(8);
  OpenAIKeyEdit.Width := OpenAIKeyPage.SurfaceWidth - ScaleX(16);
  OpenAIKeyEdit.Height := ScaleY(18);
  OpenAIKeyEdit.Text := '';
  // OpenAIKeyEdit.PasswordChar := '*'; // Uncomment to mask the input
end;

function NextButtonClick(CurPageID: Integer): Boolean;
begin
  Result := True;
  if CurPageID = OpenAIKeyPage.ID then
  begin
    if not OpenAIKeyValid(OpenAIKeyEdit.Text) then
    begin
      MsgBox('The OpenAI API Key cannot be empty and must start with ''sk-''.', mbError, MB_OK);
      Result := False;
    end;
  end;
end;

procedure CurStepChanged(CurStep: TSetupStep);
var
  FileName: string;
  APIKey: string;
  SaveResult: Boolean;
begin
  if CurStep = ssPostInstall then
  begin
    FileName := ExpandConstant('{app}\gopt-renamer.conf');
    APIKey := 'OPENAI_API_KEY=' + OpenAIKeyEdit.Text;

    // Attempt to save the API key to the file
    SaveResult := SaveStringToFile(FileName, APIKey, False);

    // Check if there was an error during saving
    if not SaveResult then
    begin
      MsgBox('Error: Unable to save the API key to the configuration file.', mbError, MB_OK);
    end;
  end;
end;
