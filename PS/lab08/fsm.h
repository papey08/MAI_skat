#ifndef FSM_H
#define FSM_H
#include <vector>
#include <map>
#include <set>

class tFSM{
public:
  typedef char tSymbol;
  typedef unsigned char tState;
  typedef std::set<tState> tStateSet;

  tFSM(){};
  void add(tState from, tSymbol c, tState to);
  void final(tState st);
  int  apply(const tSymbol* input);

  size_t size()const{return table.size();}

private:
  typedef std::map<tSymbol,tState> tTransMap;
  typedef std::vector<tTransMap> tStateTable;

  tStateTable 	table;
  tStateSet 	finals;
};
  void addstr(tFSM& fsm,
              tFSM::tState from, const tFSM::tSymbol *str,
              tFSM::tState to);
  void addrange(tFSM& fsm,
                tFSM::tState from, tFSM::tSymbol first,
                tFSM::tSymbol last, tFSM::tState to);
//------------------------------------------------------
inline void tFSM::add(tState from,tSymbol c,tState to){
  size_t sz=1+(from > to ? from : to);//1+max(from,to)
  if (sz > table.size())table.resize(sz);
  table[from][c] = to;
}

inline void tFSM::final(tState st){finals.insert(st);}

inline int tFSM::apply(const tSymbol* input){
  if(table.empty()) return 0;
  tState state=0;
  int accepted=0;

  while (*input){
    tTransMap::iterator iter;
    tTransMap &trans=table[state];

    if ((iter=trans.find(*input))==
                     trans.end()) break;

    state = iter->second;
    ++accepted;
    ++input;
  }
  return(finals.count(state)==0)? 0 : accepted;
}

  inline void addstr(tFSM& fsm,
              tFSM::tState from, const tFSM::tSymbol *str,
              tFSM::tState to){
   for(; *str; ++str) fsm.add(from, *str, to);
  }

  inline void addrange(tFSM& fsm,
                tFSM::tState from, tFSM::tSymbol first,
                tFSM::tSymbol last, tFSM::tState to){
   for(tFSM::tSymbol i=first; i<=last; ++i) fsm.add(from, i, to);
  }

#endif
