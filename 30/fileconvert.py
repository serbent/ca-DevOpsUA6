import os


def convert_file(input_file, output_file):
    with open(input_file, 'r') as f:
        data = f.read()
    with open(output_file, 'w') as f:
        f.write(data)
    

if __name__ == '__main__':

    for filename in os.listdir('.'):    
        if filename.endswith('.txt'):
            convert_file(filename, filename + '.out')