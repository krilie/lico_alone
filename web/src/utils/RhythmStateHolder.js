export default class RhythmStateHolder {
    interval = 500 // ms
    rhythm = [0, 0, 0, 1, 1, 1] // ...---
    currentIndex = 0
    currentState = [] // {{point:time}{}}

    PushStateOrReset = () => {
        console.log(this.currentIndex,this.currentState)
        const now = Date.now()
        if (this.currentIndex === 0) {
            this.currentIndex += 1
            this.currentState.push({point: now})
            console.log("push first "+now)
            return false;
        } else {
            const interval = now - this.currentState[this.currentIndex-1].point;
            console.log("interval first "+interval)
            if (this.rhythm[this.currentIndex - 1] === 0) {
                if (interval <= this.interval) {
                    this.currentIndex += 1
                    if (this.currentIndex === 6) {
                        this.ResetState()
                        return true;
                    }
                    this.currentState.push({point: now})
                } else {
                    this.ResetState()
                }
            } else if (this.rhythm[this.currentIndex - 1] === 1) {
                if (interval <= this.interval) {
                    this.ResetState()
                } else {
                    this.currentIndex += 1
                    if (this.currentIndex === 6) {
                        this.ResetState()
                        return true;
                    }
                    this.currentState.push({point: now})
                }
            }
            return false;
        }
    }

    ResetState = () => {
        this.currentIndex = 0
        this.currentState = []
    }
}