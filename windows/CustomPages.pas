var
  IntroPage: TWizardPage;
  OpenAIKeyPage: TWizardPage;
  OpenAIKeyEdit: TEdit;
  IntroRichEditViewer: TRichEditViewer;

procedure InitializeWizard();
var
  IntroRTF: AnsiString;
begin
  // Create the introductory page
  IntroPage := CreateCustomPage(wpWelcome, 'Welcome to gopt-renamer', 'Introduction');

  // Create and configure the rich edit viewer
  IntroRichEditViewer := TRichEditViewer.Create(WizardForm);
  IntroRichEditViewer.Parent := IntroPage.Surface;
  IntroRichEditViewer.ReadOnly := True;
  IntroRichEditViewer.ScrollBars := ssNone;
  IntroRichEditViewer.Top := ScaleY(8);
  IntroRichEditViewer.Width := IntroPage.SurfaceWidth - ScaleX(16);
  IntroRichEditViewer.Height := ScaleY(60);
  IntroRichEditViewer.TabStop := False;

  // Set the introductory rich text with hyperlink
  IntroRTF := '{\rtf1\ansi Welcome to gopt-renamer! This application allows you to use the power of AI to ' +
              'rename your images and screenshots.\par ' +
              'Please visit our GitHub repository for more information: {\field{\*\fldinst{HYPERLINK "https://github.com/robin-collins/gopt-renamer"}}{\fldrslt{https://github.com/robin-collins/gopt-renamer}}}\par ' +
              'You will need an OpenAI API key to use this application. We will ask for this in the next step and ' +
              'it will be stored locally on your machine.\par ' +
              'Click Next to continue or Cancel if you do not have an API key.}';
  IntroRichEditViewer.RTFText := IntroRTF;
  
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
