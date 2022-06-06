#ifndef SLR_H
#define SLR_H
#include "kngramma.h"

class tLR{
public:
  const tGramma& gr;
//  типы
  typedef short tState;
// конструктор
  tLR(const tGramma& agr):gr(agr){}; //создает "пустой" автомат
// функции-члены
  static tState pack(tGramma::tSymb left, tGramma::tAltind ialt){
   return -((ialt & 127)*256+left);
}

  static tGramma::tRule unpack(tState a){
   tState t = -a;
   return tGramma::tRule(t & 255, t >> 8);
}

//     добавляет одну команду (from,c)->to
 void add(tState from, tGramma::tSymb c, tState to){
  size_t sz=1+from;
  if (sz > table.size())table.resize(sz);
  table[from].insert(std::make_pair(c,to));
}

 void add(tState from, const tGramma::tSymbstrset& follow,
          tGramma::tSymb left, tGramma::tAltind ialt){
  tGramma::tSymbstr s = follow[left];
  tState to = pack(left, ialt);
  size_t sz=s.size();
  for(size_t i=0; i<sz; ++i) add(from,s[i],to);
}

  tState  go(tState from, tGramma::tSymb c){//переход
    if(table.empty() || from<0) return 0;
    tTransMap::iterator iter;
    tTransMap &trans=table[from];

    if ((iter=trans.find(c))==
                     trans.end()) return 0;// нет перехода
    return iter->second;
}

  tGramma::tSymbstr  expected_tokens(tState state){
    tGramma::tSymbstr tmp;
    if( state >= table.size()) return tmp;
    tTransMap::iterator iter;
    tTransMap &trans=table[state];
    for(iter=trans.begin(); iter != trans.end();
                                        ++iter){
      tGramma::tSymb c = iter->first;
      if(gr.terminal(c)) tmp += c;
        else break;
   }//for
    return tmp;
}

  void clear(){// очищает автомат
                          table.clear();
}

  size_t size()const{return table.size();}//выдает
//       размер (количество состояний) автомата

// представление недетерминированного конечного
//            автомата
  typedef std::multimap<tGramma::tSymb,tState> tTransMap;
  typedef std::vector<tTransMap> tStateTable;

  tStateTable 	table;  //таблица состояний
};

class tLRI{
public:
  typedef unsigned char tPoint;
  tGramma::tSymb left;
  tGramma::tAltind ialt;
  tPoint point;
  tGramma::tSymb smb;
// конструктор
  tLRI(tGramma::tSymb aleft=0, tGramma::tAltind aialt=0,
       tPoint apoint=0, tGramma::tSymb asmb=0)
       :left(aleft), ialt(aialt), point(apoint), smb(asmb){};
// функции-члены
  tLRI& first_point (tGramma gr){
   const tGramma::tSymbstr& rp= gr.rightPart(left,ialt);
   point =0;
   smb = rp[0]; 
   return *this;
}

  tLRI& next_point (tGramma gr){
   const tGramma::tSymbstr& rp= gr.rightPart(left,ialt);
   size_t sz = rp.size();
   if(point>=sz) return *this;
   ++point;
   smb = (point<sz ? rp[point] : 0); 
   return *this;
}

  bool operator<(const tLRI& y)const{
   return (smb<y.smb)
        ||((smb==y.smb)&&((left<y.left)
        ||((left==y.left)&&((ialt<y.ialt)
        ||((ialt==y.ialt)&&(point<y.point)
          )))));
}

  bool operator==(const tLRI& y)const{
   return (smb==y.smb)&&(left==y.left)&&
          (ialt==y.ialt)&&(point==y.point);
}

};//tLRI

  typedef std::vector<tLRI> tLRitems;
  void sort_items(tLRitems& items);
  void insert_item(tLRitems& items, const tLRI& item);
  void add_to_closure(const tGramma& gr,
               tLRitems& closure, tGramma::tSymb a);
  void make_closure(const tGramma& gr, tLRitems& closure);

  typedef std::vector<tLRitems> tLRkernels;
  bool operator==(const tLRitems& x, const tLRitems& y);
  size_t insert_kernel(tLRkernels& kernels, const tLRitems& kernel);

 void SLRbuild(tLR& lr);

#endif
