@echo off

REM build windows binary
echo Creating windows build
go build -o bin\geolocate.exe cmd\geolocate\main.go