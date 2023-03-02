'use strict';

document.querySelectorAll('.social-reaction-summary-item').forEach(aria => {
   const label = aria.getAttribute('aria-label');
   console.log(label);
});
