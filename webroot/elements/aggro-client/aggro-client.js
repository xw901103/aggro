/**
 *
 * Developed by Xu <xw901103@gmail.com>
 *
 */
import { html, PolymerElement } from '/static/modules/polymer/polymer/polymer-element.js';

class AggroClient extends PolymerElement {
  static get template() {
    return html``
  }

  static get properties() {
    return {
    }
  }

  constructor() {
    super();
  }
}

window.customElements.define(AggroClient.is, AggroClient);
