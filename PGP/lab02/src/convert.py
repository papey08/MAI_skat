from PIL import Image
import struct
import ctypes
import sys

def from_png_to_data(png_path, data_path):
    img = Image.open(png_path)
    (w, h) = img.size[0:2]
    pix = img.load()
    buff = ctypes.create_string_buffer(4 * w * h)
    offset = 0
    for j in range(h):
        for i in range(w):
            r = bytes((pix[i, j][0],))
            g = bytes((pix[i, j][1],))
            b = bytes((pix[i, j][2],))
            a = bytes((255,))
            struct.pack_into('cccc', buff, offset, r, g, b, a)
            offset += 4
    out = open(data_path, 'wb')
    out.write(struct.pack('ii', w, h))
    out.write(buff.raw)
    out.close()

def from_data_to_png(png_path, data_path):
    fin = open(data_path, 'rb')
    (w, h) = struct.unpack('hi', fin.read(8))
    buff = ctypes.create_string_buffer(4 * w * h)
    fin.readinto(buff)
    fin.close()
    img = Image.new('RGBA', (w, h))
    pix = img.load()
    offset = 0
    for j in range(h):
        for i in range(w):
            (r, g, b, a) = struct.unpack_from('cccc', buff, offset)
            pix[i, j] = (ord(r), ord(g), ord(b), ord(a))
            offset += 4
    img.save(png_path)

if __name__ == '__main__':
    if '--to_d' in sys.argv and '--from_d' in sys.argv:
        print('Only one flag allowed',
              'See doc with --help', sep='\n')
    elif '--help' in sys.argv:
        print('Usage of convert.py:',
              'FLAGS:',
              '--to_d in.png in.data -- create file in.data and convert file in.png to in.data',
              '--from_d out.png out.data -- create file out.png and convert file out.data to out.png',
              '--help -- see this documentation', sep='\n')
    elif '--to_d' in sys.argv and len(sys.argv) == 4:
        from_png_to_data(sys.argv[2], sys.argv[3])
    elif '--from_d' in sys.argv and len(sys.argv) == 4:
        from_data_to_png(sys.argv[2], sys.argv[3])
    else:
        print('Incorrect usage',
              'See doc with --help', sep='\n')
