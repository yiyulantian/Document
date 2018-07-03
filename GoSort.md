# Golang中Sort的内部实现（上）
Go与其他常见编程语言一样，库函数中提供了一系列实现排序的方法：  
*   针对简单类型（int，float，string）的数组切片，有函数可以直接进行排序，也有扩展性更好的对应类型定制  
*   针对Map和自定义结构，提供了接口继承，实现基础方法即可  

无论使用哪种方式，语言内部对排序的实现逻辑是唯一的，这就要求库函数不能单纯的图快，而是要确保在各种数据集下都不慢，因为谁都不知道使用者会传入什么样的数据。为了现实这个目的，库函数通常会混合使用多种排序算法来应对不同的场景。  

而Go语言，本身强类型且Map无序，这导致了需要借助库函数实现排序的场景比其他语言要多。伴随着每个版本的发布，sort包多少都有一些改动；从这些代码变更中，不难看出撰写者也在不断尝试使用更先进，更实用的算法来优化核心的排序逻辑。这里针对网上容易找到的两个版本文档（Golang中文社区翻译的中文文档1.4和最新版本1.10的英文文档），来探讨一下Sort的内部实现以及当时的代码思路。  
  
## 1.4版本的内部实现
注：以下代码片段中的注释（包括中文注释）均出自[文档](http://docscn.studygolang.com/src/sort/sort.go?s=6166:6191#L218)摘录  

这个版本中，sort库函数实现了插入排序，堆排和快排三种基础排序算法；其中快排是采用了改进过的三路切分算法：fat partitioning（[Bentley and McIlroy,"Engineering a Sort Function" SP&E November 1993.](http://pauillac.inria.fr/~maranget/X/421/09/bentley93engineering.pdf) ）。具体实现涉及的函数如下：
```golang
func Sort(data Interface)  
func quictSort(data Interface, a, b, maxDepth int)  
func doPivot(data Interface, lo, hi int) (midlo, midhi int)  
func swapRange(data Interface, a, b, n int)  
func medianOfThree(data Interface, m1, m0, m2 int)  
func heapSort(data Interface, a, b int)  
func siftDown(data Interface, lo, hi, first int)  
func insertionSort(data Interface, a, b int)  
func min(a, b int) int  
```
其中, data Interface 即是Go语言对外提供的排序接口；实现其内部方法、继承接口就可以使用Sort方法排序，这里不再详述了。下面说说各算法的内部实现和意义：  

### 插入排序  
涉及到函数有：insertionSort。代码是基础实现没什么说的～
```golang
    func insertionSort(data Interface, a, b int) {
        for i := a + 1; i < b; i++ {
            for j := i; j > a && data.Less(j, j-1); j-- {
                data.Swap(j, j-1)
            }
        }
    }
```
### 堆排
涉及到函数有：siftDown，heapSort。
*   siftDown：用来维护堆性质的子函数；可以理解为除根节点（root：lo）以外其他子节点已经实现了堆性质；由于根节点的加入，需要重新维护堆性质。该函数就是用来维护数据集堆性质的
*   heapSort：堆排入口，内部调用siftDown；注释中最大元素和从大到小的顺序的相关描述，并不准确；因为元素的大小以及升序降序都和less()函数的实现有关；这里只是以数值的大小、升序排序来简单说明一下，便于理解。组织大顶堆，根结点弹出放到队尾，实现升序效果

```golang
    // siftDown implements the heap property on data[lo, hi).
    // first is an offset into the array where the root of the heap lies.
    
    // siftDown 为 data[lo, hi) 实现堆的性质。
    // first 为堆的根节点在数组中的偏移量。
    func siftDown(data Interface, lo, hi, first int) {
        root := lo
        for {
            child := 2*root + 1
            if child >= hi {
                break
            }
            if child+1 < hi && data.Less(first+child, first+child+1) {
                child++
            }
            if !data.Less(first+root, first+child) {
                return
            }
            data.Swap(first+root, first+child)
            root = chil
        }
    }


    func heapSort(data Interface, a, b int) {
        first := a
        lo := 0
        hi := b - a
    
        // Build heap with greatest element at top.
        // 以最大元素为顶建堆。
        for i := (hi - 1) / 2; i >= 0; i-- {
            siftDown(data, i, hi, first)
        }
    
        // Pop elements, largest first, into end of data.
        // 弹出元素，从大到小的顺序，从后向前依次追加到数组 data。
        for i := hi - 1; i >= 0; i-- {
            data.Swap(first, first+i)
            siftDown(data, lo, i, first)
        }
    }    
```
### 快排
涉及到函数有：quickSort，doPivot，swapRange，medianOfThree，min。
*   min：两个整数取较小的那个，没什么好说的
*   medianOfThree：用于选定合适的中轴数；对于快排来说，中轴数的好坏直接影响到排序的效率。该函数负责从数据集候选的三个位置中选择一个中间值,并交换到m1位置，使得：data.Less(m1, m0)==false && data.Less(m2, m1)==false
*   swapRange：批量交换两段位置上的值；为了支持多个相等的中轴数集中调换位置
*   doPivot：将一段数据分成三段（排在中轴数之前的数，中轴数，排在中轴数之后的数），钉死中轴数的位置，返回中轴数的开始、结束位置；内部调用了medianOfThree、swapRange、min函数。在寻找中轴数时，先求出这段数据的中间位置（这里为了避免两数直接相加溢出，使用了先减后加的方式，这是库函数应有的严谨写法）；当数据段长度大于40时，为了更均匀的寻找中轴数，先将数据分8段（ s := (hi - lo) / 8 ），包括头尾在内可以得到九个数，先做一次medianOfThree操作，得到三个数后再做一次medianOfThree，最终得到中轴数。具体为什么这么做，这么做有什么优点可以根据文档注释查看一下相关论文。
*   quickSort：快排入口；这里的快排已经不是单纯的快排算法实现了，而是快排，堆排，插入排序的混合；根据预定的逻辑，在适当的时候使用合适的排序算法；也可以说这是个中控函数。内部调用了insertionSort、heapSort、doPivot；快排的实现用到了递归，为了避免递归太深影响性能，函数多传入了一个maxDepth参数来限制最大的递归深度。根据逻辑，当数据段长度小于等于7时，会切换到插入排序；当长度大于7时，先使用快排，当递归深度到达maxDepth时，切换成堆排（堆排不是递归实现的）。

```golang

    func min(a, b int) int {
    	if a < b {
    		return a
    	}
    	return b
    }

    // medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
    
    // medianOfThree 将 data[m1]、data[m2] 和 data[m3] 三个值的中值交换到 data[m1]。
    func medianOfThree(data Interface, m1, m0, m2 int) {
        // sort 3 elements
        if data.Less(m1, m0) {
            data.Swap(m1, m0)
        }
        // data[m0] <= data[m1]
        if data.Less(m2, m1) {
            data.Swap(m2, m1)
            // data[m0] <= data[m2] && data[m1] < data[m2]
            if data.Less(m1, m0) {
                data.Swap(m1, m0)
            }
        }
        // now data[m0] <= data[m1] <= data[m2]
        // 现在 data[m0] <= data[m1] <= data[m2]
    }

    func swapRange(data Interface, a, b, n int) {
        for i := 0; i < n; i++ {
            data.Swap(a+i, b+i)
        }
    }


    func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
        m := lo + (hi-lo)/2 // Written like this to avoid integer overflow. // 这样写避免整数溢出。
        if hi-lo > 40 {
            // Tukey's ``Ninther,'' median of three medians of three.
            // Tukey的“Ninther”算法，分别求三组值的中值，再求这三个值的中值。
            s := (hi - lo) / 8
            medianOfThree(data, lo, lo+s, lo+2*s)
            medianOfThree(data, m, m-s, m+s)
            medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
        }
        medianOfThree(data, lo, m, hi-1)
    
        // Invariants are:
        //    data[lo] = pivot (set up by ChoosePivot)
        //    data[lo <= i < a] = pivot
        //    data[a <= i < b] < pivot
        //    data[b <= i < c] is unexamined
        //    data[c <= i < d] > pivot
        //    data[d <= i < hi] = pivot
        //
        // Once b meets c, can swap the "= pivot" sections
        // into the middle of the slice.
    
        // 算法不变式为：
        //    data[lo] = pivot (由 ChoosePivot 决定)
        //    data[lo <= i < a] = pivot
        //    data[a <= i < b] < pivot
        //    data[b <= i < c] 未经检查
        //    data[c <= i < d] > pivot
        //    data[d <= i < hi] = pivot
        //
        // 一旦 b 与 c 相遇，就将“= pivot”的这部分交换到切片中间。
        pivot := lo
        a, b, c, d := lo+1, lo+1, hi, hi
        for {
            for b < c {
                if data.Less(b, pivot) { // data[b] < pivot
                    b++
                } else if !data.Less(pivot, b) { // data[b] = pivot
                    data.Swap(a, b)
                    a++
                    b++
                } else {
                    break
                }
            }
            for b < c {
                if data.Less(pivot, c-1) { // data[c-1] > pivot
                    c--
                } else if !data.Less(c-1, pivot) { // data[c-1] = pivot
                    data.Swap(c-1, d-1)
                    c--
                    d--
                } else {
                    break
                }
            }
            if b >= c {
                break
            }
            // data[b] > pivot; data[c-1] < pivot
            data.Swap(b, c-1)
            b++
            c--
        }
    
        n := min(b-a, a-lo)
        swapRange(data, lo, b-n, n)
    
        n = min(hi-d, d-c)
        swapRange(data, c, hi-n, n)
    
        return lo + b - a, hi - (d - c)
    }


    func quickSort(data Interface, a, b, maxDepth int) {
        for b-a > 7 {
            if maxDepth == 0 {
                heapSort(data, a, b)
                return
            }
            maxDepth--
            mlo, mhi := doPivot(data, a, b)
            // Avoiding recursion on the larger subproblem guarantees
            // a stack depth of at most lg(b-a).
            // 避免大量子问题的递归，以保证栈的深度在 lg(b-a) 以内。
            if mlo-a < b-mhi {
                quickSort(data, a, mlo, maxDepth)
                a = mhi // i.e., quickSort(data, mhi, b) // 例如 quickSort(data, mhi, b)
            } else {
                quickSort(data, mhi, b, maxDepth)
                b = mlo // i.e., quickSort(data, a, mlo) // 例如 quickSort(data, a, mlo)
            }
        }
        if b-a > 1 {
            insertionSort(data, a, b)
        }
    }

```  

### 排序入口：Sort函数
这个函数就是在代码中常用的Sort方法了；是数据集排序的入口。之前说过快排是递归实现的，为了避免无限递归带来的问题，需要传入一个允许的最深深度，超过就切换成非递归的堆排序；而这里就说明了这个深度取多少合适。根据函数的逻辑可以看出，允许递归的最大深度是2 * ceil(log(n+1))（个人认为代码中注释应该写错了，>>= 1的效果应该是log而不是lg）。根据二叉树的性质，ceil(log(n+1))是n个元素可以组成的完全二叉树的高，换句话说也是快排最优、最快的情况：每次中轴数都可以平分数据集；而2 * ceil(log(n+1))是完全二叉树非叶子结点的个数，换句话说也是快排最优情况对应的递归深度。
```golang
    // Sort sorts data.
    // It makes one call to data.Len to determine n, and O(n*log(n)) calls to
    // data.Less and data.Swap. The sort is not guaranteed to be stable.
    
    // Sort 对 data 进行排序。
    // 它调用一次 data.Len 来决定排序的长度 n，调用 data.Less 和 data.Swap 的开销为
    // O(n*log(n))。此排序为不稳定排序。
    func Sort(data Interface) {
        // Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
        // 若深度达到 2*ceil(lg(n+1)) 就切换到堆排序算法。
        n := data.Len()
        maxDepth := 0
        for i := n; i > 0; i >>= 1 {
            maxDepth++
        }
        maxDepth *= 2
        quickSort(data, 0, n, maxDepth)
    }
```
### 版本总结
该版本的排序库函数，实现了插入、堆排、快排三种基础排序算法，并根据规则不断切换排序算法，以确保时间、空间上的资源平衡：数量少用插入，数量多用快排，递归深了换堆排。除了一些已经得到论证的改进，三个算法的实现代码与经典代码出入不大。
##  1.10版本的内部实现