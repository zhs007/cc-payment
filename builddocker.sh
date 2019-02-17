docker build -t cc-payment .

if [ ! -d "logs" ]; then
    mkdir logs
fi