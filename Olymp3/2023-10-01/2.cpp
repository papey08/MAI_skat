#include <iostream>
#include <fstream>
#include <string>

std::string getVersionAfterLetterPatch(std::string curVersion) {
	if (curVersion.size() == 4) { // áóêâåííîãî ïàò÷à íåò
		return curVersion + "b";
	}
	else {
		int letterCode = curVersion[4];
		
		++letterCode;
		char letter = (char)letterCode;
		//std::cout << letter << std::endl;
		curVersion[4] = letter;
		std::string newVersion = curVersion;
		return newVersion;
	}
}

std::string getVersionAfterMinor(std::string curVersion) {
	int unitsCode = curVersion[3];
	int tensCode = curVersion[2];
	std::string newVersion;
	if (unitsCode != 57) {
		++unitsCode;
		char newUnits = (char)unitsCode;
		curVersion[3] = newUnits;
		newVersion = curVersion.substr(0, 4);
	}
	else {
		char newUnits = '0';
		++tensCode;
		char newTens = (char)tensCode;
		curVersion[3] = newUnits;
		curVersion[2] = newTens;
		newVersion = curVersion.substr(0, 4);
	}
	return newVersion;
}

std::string getVersionAfterMajor(std::string curVersion) {
	int majorCode = curVersion[0];
	++majorCode;
	char newMajor = (char)majorCode;
	curVersion[0] = newMajor;
	curVersion[2] = '0';
	curVersion[3] = '0';
	std::string newVersion = curVersion.substr(0, 4);
	return newVersion;
}

int main() {
	std::ifstream in("input.txt");
	std::string curVersion;
	in >> curVersion;
	in.close();
	std::ofstream out("output.txt");
	out << getVersionAfterLetterPatch(curVersion) << std::endl;
	out << getVersionAfterMinor(curVersion) << std::endl;
	out << getVersionAfterMajor(curVersion) << std::endl;
	out.close();
	return 0;
}
