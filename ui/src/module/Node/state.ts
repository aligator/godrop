import {Observable} from "../../plainJSX";
import {FileNode, GetNodeDocument, GetNodeQuery, GetNodeQueryVariables} from "../../api/types";
import client from "../../api/apollo";

export interface FileNodeState {
    fileNode: Observable<FileNode | undefined>
    currentPath: Observable<string | undefined>
    load(path: string): void
    destroy(): void
}

export function useFileNodeState(initialPath: string): FileNodeState {
    const fileNode = new Observable<FileNode | undefined>(undefined)
    const currentPath = new Observable<string | undefined>(undefined)

    const load = (path: string) => {
        console.log("load")

        const variables: GetNodeQueryVariables = {
            path
        }

        void client.query<GetNodeQuery>({
            query: GetNodeDocument,
            variables,
        }).then((res) => {
            currentPath.value = path
            fileNode.value = res.data.getNode
        })
    }

    load(initialPath)

    return {
        fileNode,
        load,
        currentPath,
        destroy: () => {
            fileNode.destroy()
            currentPath.destroy()
        }
    }
}

