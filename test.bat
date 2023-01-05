@echo off
@REM mkdir output
@REM mkdir csv
@REM
@REM echo Voici la liste des fichiers dans le repertoire "output" :
@REM
@REM cd output
@REM set count=0
@REM for /r %%f in (*) do (
@REM   set /a count+=1
@REM   echo [%count%]: %%f
@REM )
@REM
@REM dir /b *
@REM
@REM set /p fileNum=Veuillez saisir le * du fichier à lancer :
@REM set /a fileNum-=1
@REM set index=0
@REM for /r %%f in (*) do (
@REM   if %index%==%fileNum% (
@REM     start "" "%%f"
@REM     goto end
@REM   )
@REM   set /a index+=1
@REM )
@REM :end
@REM
@REM pause


:afficher_fichiers

set count=0
cd output
for %%A IN (*) DO (
    set /a count+=1
    echo %count%
    )
echo.
echo Choisissez un fichier à exécuter en saisissant son numéro :

set /p numero=

start %repertoire%\%numero%

echo.
echo Le fichier a été exécuté.
echo.

goto afficher_fichiers
