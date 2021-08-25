import PlainJSX, {PlainJSXElement, useRef} from "../../plainJSX";

import {FileNodeState} from "./state";

interface Props {
    fileNodeState: FileNodeState
}

export default function ChildList({fileNodeState}: Props): PlainJSXElement {
    const ref = useRef<HTMLUListElement>()
    fileNodeState.fileNode.listen((v) => {
        if (!v) {
            return
        }
        const listItems = (
            <>
                {(v.children || []).map((child) => {
                    return <li onclick={() => fileNodeState.load(`${fileNodeState.currentPath?.value || ""}/${child.name}`)}>{child.name}</li>
                })}
            </>
        )
        ref.update((list) => {
            listItems.setChildOf(list)
        })
    })

    return (
        <ul id={ref.id}></ul>
    )
}