import { useAsync } from 'react-use';

import { config } from '@grafana/runtime';

import { preloadPlugins } from '../pluginPreloader';

import { getAppPluginConfigs } from './utils';

export function useLoadAppPlugins(pluginIds: string[] = []): { isLoading: boolean } {
  const { loading: isLoading } = useAsync(async () => {
    const isEnabled = config.featureToggles.appPluginLazyLoading;
    const appConfigs = getAppPluginConfigs(pluginIds);

    if (!isEnabled || !appConfigs.length) {
      return;
    }

    await preloadPlugins(appConfigs);
  });

  return { isLoading };
}
