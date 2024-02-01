@echo off
SETLOCAL ENABLEDELAYEDEXPANSION

echo Starting script execution...

REM Set the name of the input file
SET input_video=in_action.mp4
SET output_gif=in_action.gif
SET palette_file=palette.png

echo Generating palette...

REM Generate the palette from the first half of the video
ffmpeg -i "%input_video%" -vf "fps=10,scale=-1:1080:flags=lanczos,palettegen" -t 50 -y "%palette_file%"

echo Creating GIF...

REM Create the GIF using the custom palette
ffmpeg -i "%input_video%" -i "%palette_file%" -filter_complex "fps=10,scale=-1:1080:flags=lanczos[x];[x][1:v]paletteuse" -f gif "%output_gif%"

echo Process completed.
pause
