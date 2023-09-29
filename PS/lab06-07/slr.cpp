//                 slr.cpp
#include <iostream>
#include <iomanip>
#include "slr.h"

  void sort_items(tLRitems& items){
  size_t n=items.size();
   if(n<2)return;
   --n;
   for (size_t j=0; j<n; ++j){
     for(size_t i=0; i<n-j; ++i){
       if(items[i+1]<items[i]){
         tLRI t=items[i];
         items[i]=items[i+1];
         items[i+1]=t;
      }//if
    }//for i
  }//for j
}

  void insert_item(tLRitems& items, const tLRI& item){
  size_t n=items.size();
   for (size_t i=0; i<n; ++i){
    if(items[i]==item) return;
   }//for i
   items.push_back(item);
}

  void add_to_closure(const tGramma& gr,
               tLRitems& closure, tGramma::tSymb a){
  size_t n=gr.altCount(a);
   for (size_t i=0; i<n; ++i){
    insert_item(closure, tLRI(a, i).first_point(gr));
   }//for i
}

  void make_closure(const tGramma& gr, tLRitems& closure){
   for (size_t i=0; i<closure.size(); ++i){
    tLRI& item = closure[i];
    if(!gr.terminal(item.smb))
              add_to_closure(gr, closure, item.smb);
  }//for i
}

  bool operator==(const tLRitems& x, const tLRitems& y){
   if(x.size()!=y.size()) return false;
   for (size_t i=0; i<x.size(); ++i){
    if(!(x[i]==y[i])) return false;
  }//for i
  return true;
}

  size_t insert_kernel(tLRkernels& kernels, const tLRitems& kernel){
  size_t n=kernels.size();
   for (size_t i=0; i<n; ++i){
    if(kernels[i] == kernel) return i;
   }//for i
   kernels.push_back(kernel);
  return n;
}


 void SLRbuild(tLR& lr){
   const tGramma& gr(lr.gr);
   tGramma::tSymbstrset follow;
   gr.makeFollow(follow);

  lr.clear();
 tLRI item;
 tLRitems closure;
 tLRitems kernel;
 tLRkernels kernels;

 tGramma::tSymb smb=gr.getStart();
 for(size_t i=0; i<gr.altCount(smb); ++i){
  item = tLRI(smb,i).first_point(gr);
  insert_item(kernel,item);
 }
 kernels.push_back(kernel);

 for(size_t from=0; from<kernels.size(); ++from){
  closure.clear();
  size_t to;

  for(size_t i=0; i<kernels[from].size(); ++i){
   tLRI& ritem = kernels[from][i];
   if(ritem.smb)
                closure.push_back(ritem);
     else lr.add(from, follow, ritem.left, ritem.ialt);
  }//for i

  if(!closure.empty()){
   make_closure(gr, closure);
   sort_items(closure);
   item = closure[0];
   smb =item.smb;
   item.next_point(gr);
   kernel.clear();
   kernel.push_back(item);
   size_t i=1;
   for(; i<closure.size(); ++i){
     item = closure[i];
     if(smb == item.smb){
       item.next_point(gr);
       kernel.push_back(item);
       continue;
     }// if smb
     to = insert_kernel(kernels,kernel);
     lr.add(from, smb, to);
     smb =item.smb;
     item.next_point(gr);
     kernel.clear();
     kernel.push_back(item);
   }//for i 
   to = insert_kernel(kernels,kernel);
   lr.add(from, smb, to);
  }//if !closure
 }// for from

// out_kernels(cout, gr, kernels);
}
