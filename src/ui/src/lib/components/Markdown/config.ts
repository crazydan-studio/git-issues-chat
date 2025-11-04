import { Carta } from 'carta-md';
import DOMPurify from 'isomorphic-dompurify';
import { math } from '@cartamd/plugin-math';
import { attachment } from '@cartamd/plugin-attachment';
import { emoji } from '@cartamd/plugin-emoji';
import { slash } from '@cartamd/plugin-slash';
import { code } from '@cartamd/plugin-code';
import 'carta-md/default.css';
import '@cartamd/plugin-code/default.css';
import '@cartamd/plugin-emoji/default.css';
import '@cartamd/plugin-slash/default.css';
import '@cartamd/plugin-attachment/default.css';
import 'katex/dist/katex.min.css';

export default function createCarta() {
  return new Carta({
    sanitizer: DOMPurify.sanitize,
    extensions: [
      math(),
      attachment({
        async upload() {
          return 'some-url-from-server.xyz';
        }
      }),
      emoji(),
      slash(),
      code({theme: ''})
    ]
  });
}