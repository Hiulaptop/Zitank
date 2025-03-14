import {
  DecoratorNode,
  DOMConversionMap,
  DOMConversionOutput,
  DOMExportOutput,
  LexicalNode,
  NodeKey,
} from "lexical";
import { JSX } from "react";

export const $createYoutubeNode = ({ id }: { id: string }) => {
  return new YoutubeNode({ id });
};

const ID_ATTR = "data-lexical-youtube";

const convertYoutubeElement = (domNode: HTMLElement): DOMConversionOutput | null => {
  const id = domNode?.getAttribute(ID_ATTR);
  if (!id) return null;
  return { node: $createYoutubeNode({ id }) };
};

const HEIGHT = "315px";
const WIDTH = "560px";

const getYoutubeLink = (id: string) => `https://www.youtube-nocookie.com/embed/${id}`;

export class YoutubeNode extends DecoratorNode<JSX.Element> {
  __id: string;

  constructor({ id, key }: { id: string; key?: NodeKey }) {
      super(key);
      this.__id = id;
  }

  static getType(): string {
      return "youtube";
  }

  static clone(_node: YoutubeNode): YoutubeNode {
      return new YoutubeNode({
          id: _node.__id,
          key: _node.__key, // Đảm bảo node được clone với key chính xác
      });
  }

  decorate(): JSX.Element {
      return (
          <iframe
              height={HEIGHT}
              width={WIDTH}
              src={getYoutubeLink(this.__id)}
              frameBorder="0"
              allowFullScreen
          />
      );
  }

  createDOM(): HTMLElement {
      const div = document.createElement("div");
      div.setAttribute(ID_ATTR, this.__id);
      return div;
  }

  updateDOM(): boolean {
      return false; // React sẽ quản lý cập nhật
  }

  exportDOM(): DOMExportOutput {
      const iframe = document.createElement("iframe");
      iframe.setAttribute(ID_ATTR, this.__id);
      iframe.setAttribute("height", HEIGHT);
      iframe.setAttribute("width", WIDTH);
      iframe.setAttribute("src", getYoutubeLink(this.__id));
      iframe.setAttribute("frameBorder", "0");
      iframe.setAttribute("allowFullScreen", "true");

      return { element: iframe };
  }

  static importDOM(): DOMConversionMap | null {
      return {
          iframe: (node: Node) => ({ conversion: convertYoutubeElement, priority: 0 }),
      };
  }
}
