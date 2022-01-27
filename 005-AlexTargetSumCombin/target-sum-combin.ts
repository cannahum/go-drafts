function getTargetSumCombin(arr: number[], target: number): number[][] {
    const result: number[][] = [];
    if (arr.length === 0 || target === 0) {
        return result;
    }
    // console.log(arr);

    for (let i = arr.length - 1; i >= 0; i--) {
        const nextMax = arr[i];
        if (nextMax > target) {
            continue;
        }
        if (nextMax === target) {
            result.push([nextMax]);
            continue;
        }
        // this works! why? - if you do arr.slice(0, i) then it doesn't
        const remainingArr = arr.slice(0, i-1);
        console.log(remainingArr);
        for (let y of getTargetSumCombin(remainingArr, target-nextMax)) {
            y.push(nextMax);
            result.push(y);
        }
        // getTargetSumCombin(arr.slice(0, i-1), target - nextMax).forEach((combin: number[]) => {
        //    combin.push(nextMax);
        //    result.push(combin);
        // });
    }
    return result;
}

const result = getTargetSumCombin([1, 2, 3, 4, 5, 6, 7], 7);
// console.log(result);