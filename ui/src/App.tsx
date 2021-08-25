import PlainJSX, {PlainJSXElement, useRef} from "./plainJSX/index";
import {ChildList} from "./module/Node";
import {FileNodeState, useFileNodeState} from "./module/Node/state";

export default function App(): PlainJSXElement {
    const divRef = useRef()
    let lastList: PlainJSXElement
    let lastState: FileNodeState
    return (
        <>
            <button onclick={() => {
                divRef.update((el) => {
                    if (lastList) {
                        lastList.delete()
                    }

                    if (lastState) {
                        lastState.destroy()
                    }

                    const fileNodeState = useFileNodeState("")
                    const newList = <ChildList fileNodeState={fileNodeState} />
                    lastState = fileNodeState
                    lastList = newList

                    newList.setChildOf(el)
                })
            }}>NEW</button>
            <div id={divRef.id}>

            </div>
        </>
    )
}