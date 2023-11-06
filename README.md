# Scalable Image Server

This is project 1 part 1 for the course "Operating Systems and Distributed Systems" in IIIS, Tsinghua University, 2023 Fall.

We implement a scalable image server that can resize images concurrently with channel-based synchronization in Go. The source image can be downloaded from http://cs231n.stanford.edu/tiny-imagenet-200.zip.

## Usage
`make run`: download the dataset, do the experiment and generate report automatically. 

`make clean`: delete all intermediate files including the resized image output and experiment results.

## Result
Our experiment tests both throughput and latency.

For throughput, the result shows that capacity is not the bottleneck but number of threads.

As for latency, we test against multiple capacities and different ratios of enqueue thread number over dequeue. Our result shows that capacity is not the bottleneck while dequeue latency is significantly larger than enqueue latency. Also, the higher the ratio is, the lower dequeue latency is. This strongly shows that enqueue thread number is truly the bottleneck, which is really interesting.

More details can be found in `report.pdf`.