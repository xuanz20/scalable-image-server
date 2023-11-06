import os
import matplotlib.pyplot as plt
import numpy as np

folder_path = 'figure'

# create figure folder
# if already exists, clear it
if os.path.exists(folder_path):
    for file_name in os.listdir(folder_path):
        file_path = os.path.join(folder_path, file_name)
        try:
            if os.path.isfile(file_path):
                os.remove(file_path)
        except Exception as e:
            print(f"error: {e}")
else:
    try:
        os.mkdir(folder_path)
    except Exception as e:
        print(f"error: {e}")

# plot throughput figure
with open('throughput/throughput.txt', 'r') as file:
    lines = file.readlines()

first_line = float(lines[0])
data = [list(map(float, line.strip().split(','))) for line in lines[1:]]

x_labels = [f'{int(x[0])}, {int(x[1])}' for x in data]
third_values = [x[2] for x in data]

plt.figure(figsize=(10, 6))
plt.bar(x_labels, third_values)

plt.title(f'Bar Chart')
plt.xlabel('thread number and capacity')
plt.ylabel('throughput')

plt.axhline(y=first_line, color='red', linestyle='--', label=f'Baseline: {first_line}')

for i, (x, y) in enumerate(zip(x_labels, third_values)):
    plt.text(i, y + 0.01, x, ha='center', va='bottom', rotation=45)

plt.legend()
plt.savefig('figure/throughput.png')

# plot latency figure
def calculate_cdf(data):
    sorted_data = np.sort(data)
    n = len(data)
    cdf = np.arange(1, n+1) / n
    return sorted_data, cdf

enqueue_path = 'latency/enqueue'
dequeue_path = 'latency/dequeue'

fig, ax = plt.subplots()

for txt_name in os.listdir(enqueue_path):
    _, x, y, z = txt_name.split('_')
    txt_path = os.path.join(enqueue_path, txt_name)
    with open(txt_path, 'r') as file:
        data = [int(line.strip()) for line in file.readlines()]
    
    # data = [1e-9 if value == 0 else value for value in data]
    sorted_data, cdf = calculate_cdf(data)
    sorted_data = np.log10(sorted_data)
    ax.plot(sorted_data, cdf, label=f'enqueue_latency_{x}_{y}_{z}')

ax.set_title('Enqueue Latency CDF')
ax.set_xlabel('Latency')
ax.set_ylabel('CDF')
ax.legend()

plt.savefig('figure/enqueue_latency_cdf.png')
plt.close()

fig, ax = plt.subplots()

for txt_name in os.listdir(dequeue_path):
    _, x, y, z = txt_name.split('_')
    txt_path = os.path.join(dequeue_path, txt_name)
    with open(txt_path, 'r') as file:
        data = [int(line.strip()) for line in file.readlines()]
    
    # data = [1e-9 if value == 0 else value for value in data]
    sorted_data, cdf = calculate_cdf(data)
    sorted_data = np.log10(sorted_data)
    ax.plot(sorted_data, cdf, label=f'dequeue_latency_{x}_{y}_{z}')

ax.set_title('Dequeue Latency CDF')
ax.set_xlabel('Latency')
ax.set_ylabel('CDF')
ax.legend()

plt.savefig('figure/dequeue_latency_cdf.png')
plt.close()
