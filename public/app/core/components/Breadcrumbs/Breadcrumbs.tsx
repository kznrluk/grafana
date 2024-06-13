import { css, cx } from '@emotion/css';
import React from 'react';

import { GrafanaTheme2 } from '@grafana/data';
import { useStyles2 } from '@grafana/ui';

import { BreadcrumbItem } from './BreadcrumbItem';
import { Breadcrumb } from './types';

export interface Props {
  breadcrumbs: Breadcrumb[];
  className?: string;
  isMenuDocked?: boolean;
}

export function Breadcrumbs({ breadcrumbs, className, isMenuDocked }: Props) {
  const styles = useStyles2(getStyles);

  return (
    <nav aria-label="Breadcrumbs" className={cx(className, isMenuDocked && styles.isDocked)}>
      <ol className={styles.breadcrumbs}>
        {breadcrumbs.map((breadcrumb, index) => (
          <BreadcrumbItem
            {...breadcrumb}
            isCurrent={index === breadcrumbs.length - 1}
            key={index}
            index={index}
            flexGrow={getFlexGrow(index, breadcrumbs.length)}
          />
        ))}
      </ol>
    </nav>
  );
}

function getFlexGrow(index: number, length: number) {
  if (length < 5 && index > 0 && index < length - 2) {
    return 4;
  }

  if (length > 6 && index > 1 && index < length - 3) {
    return 4;
  }

  return 10;
}

const getStyles = (theme: GrafanaTheme2) => {
  return {
    breadcrumbs: css({
      display: 'flex',
      alignItems: 'center',
      flexWrap: 'nowrap',
      overflow: 'hidden',
    }),
    isDocked: css({
      paddingLeft: theme.spacing(3),
    }),
  };
};
