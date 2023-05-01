def compare(file1, file2, n):
    with open(file1) as file1, open(file2) as file2:
        text1 = file1.read()
        text2 = file2.read()
        text1 = text1[:min(len(text1), len(text2), n)]
        text2 = text2[:min(len(text1), len(text2), n)]
        
        count = 0
        for i in range(min(len(text1), len(text2))):
            if text1[i] == text2[i] and not text1[i].isspace():
                count += 1
        return count / len(text1)
    