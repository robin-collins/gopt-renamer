[Setup]
AppName=gopt-renamer
AppVersion=0.1.0
DefaultDirName={pf}\gopt-renamer
OutputDir=..\
OutputBaseFilename=gopt-renamer-installer
Compression=lzma2
SolidCompression=yes
SetupIconFile="{#GetEnv('GITHUB_WORKSPACE')}\windows\icon.ico"
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
; New entry for setting the icon
Root: HKCR; Subkey: "SystemFileAssociations\image\shell\gopt-renamer\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: """{app}\gopt-renamer.exe"",0"


[Code]
[Code]
var
  IntroPage: TWizardPage;
  OpenAIKeyPage: TWizardPage;
  OpenAIKeyEdit: TEdit;
  IntroLabel: TLabel;

procedure InitializeWizard();
var
  IntroText: string;
begin
  // Create the introductory page
  IntroPage := CreateCustomPage(wpWelcome, 'Welcome to gopt-renamer', 'Introduction');
  
  // Create and configure the introductory label
  IntroLabel := TLabel.Create(WizardForm);
  IntroLabel.Parent := IntroPage.Surface;
  IntroLabel.Top := ScaleY(8);
  IntroLabel.Width := IntroPage.SurfaceWidth - ScaleX(16);
  IntroLabel.Height := ScaleY(60);
  IntroLabel.AutoSize := False;
  IntroLabel.WordWrap := True;
  
  // Set the introductory text
  IntroText := 'Welcome to gopt-renamer! This application allows you to use the power of AI to rename your images and screenshots.' + #13#10 +
               'Please visit our GitHub repository for more information: [link]https://github.com/robin-collins/gopt-renamer[/link]' + #13#10 +
               'You will need an OpenAI API key to use this application. We will ask for this in the next step and it will be stored locally on your machine.' + #13#10 +
               'Click Next to continue or Cancel if you do not have an API key.';
  IntroLabel.Caption := IntroText;
  
  // Create the OpenAI API key page
  OpenAIKeyPage := CreateCustomPage(IntroPage.ID, 'OpenAI API Key', 'Please enter your OpenAI API Key (it should start with ''sk-'')');
  OpenAIKeyEdit := TEdit.Create(WizardForm);
  OpenAIKeyEdit.Parent := OpenAIKeyPage.Surface;
  OpenAIKeyEdit.Top := ScaleY(8);
  OpenAIKeyEdit.Width := OpenAIKeyPage.SurfaceWidth - ScaleX(16);
  OpenAIKeyEdit.Height := ScaleY(18);
  OpenAIKeyEdit.Text := '';
  // OpenAIKeyEdit.PasswordChar := '*'; // Uncomment to mask the input
end;

function OpenAIKeyValid(Key: string): Boolean;
begin
  // Check if the key is not empty and starts with 'sk-'
  Result := (Length(Key) > 0) and (Copy(Key, 1, 3) = 'sk-');
end;

function NextButtonClick(CurPageID: Integer): Boolean;
begin
  Result := True;
  if CurPageID = IntroPage.ID then
  begin
    // Additional logic if needed when moving from the IntroPage to the OpenAIKeyPage
  end
  else if CurPageID = OpenAIKeyPage.ID then
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

