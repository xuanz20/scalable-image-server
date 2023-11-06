#!/bin/bash

if [ ! -d "tiny-imagenet-200" ]; then
    echo "Downloading tiny-imagenet-200.zip..."
    wget http://cs231n.stanford.edu/tiny-imagenet-200.zip
    
    echo "Unzipping tiny-imagenet-200.zip..."
    unzip tiny-imagenet-200.zip
    
    rm tiny-imagenet-200.zip
else
    echo "tiny-imagenet-200 folder already exists."
fi
