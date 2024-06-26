/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import ActivityComponent from '../activity';

export default class SyncComponent extends ActivityComponent {
  title = 'Secrets sync usage';
  description =
    'This data can be used to understand how many secrets sync clients have been used for this date range. A secret with a configured sync destination would qualify as a unique and active client.';
}
