@echo off
@setlocal
set CURRENT_PATH=%~dp0
if "%PHP_COMMAND%" == "" set PHP_COMMAND=php.exe
"%PHP_COMMAND%" "%CURRENT_PATH%get_random_gravatar.php" %*
@endlocal