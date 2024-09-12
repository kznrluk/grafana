import { css } from '@emotion/css';
import * as React from 'react';

import { useStyles2 } from '@grafana/ui';

import { DataSourceInformation } from '../home/Insights';

import { DataSourcesInfo } from './DataSourcesInfo';

export function SectionSubheader({
  children,
  datasources,
}: React.PropsWithChildren<{ datasources?: DataSourceInformation[] }>) {
  const styles = useStyles2(getStyles);

  return (
    <div className={styles.container}>
      {children}
      {datasources && <DataSourcesInfo datasources={datasources} />}
    </div>
  );
}

const getStyles = () => ({
  container: css({
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
  }),
});
