export interface CODEProps {
  space?: string
  repoName?: string
  repoPath?: string
  gitRef?: string
  resourcePath?: string
  commitRef?: string
  branch?: string
  diffRefs?: string
  pullRequestId?: string
  pullRequestSection?: string
  webhookId?: string
}

export interface CODEQueryProps {
  query?: string
}

export const pathProps: Readonly<Omit<Required<CODEProps>, 'repoPath' | 'branch'>> = {
  space: ':space',
  repoName: ':repoName',
  gitRef: ':gitRef*',
  resourcePath: ':resourcePath*',
  commitRef: ':commitRef*',
  diffRefs: ':diffRefs*',
  pullRequestId: ':pullRequestId',
  pullRequestSection: ':pullRequestSection*',
  webhookId: ':webhookId'
}

export interface CODERoutes {
  toSignIn: () => string
  toSignUp: () => string
  toCODERepositories: (args: Required<Pick<CODEProps, 'space'>>) => string
  toCODERepository: (args: RequiredField<Pick<CODEProps, 'repoPath' | 'gitRef' | 'resourcePath'>, 'repoPath'>) => string
  toCODEFileEdit: (args: Required<Pick<CODEProps, 'repoPath' | 'gitRef' | 'resourcePath'>>) => string
  toCODECommits: (args: Required<Pick<CODEProps, 'repoPath' | 'commitRef'>>) => string
  toCODEPullRequests: (args: Required<Pick<CODEProps, 'repoPath'>>) => string
  toCODEPullRequest: (
    args: RequiredField<
      Pick<CODEProps, 'repoPath' | 'pullRequestId' | 'pullRequestSection'>,
      'repoPath' | 'pullRequestId'
    >
  ) => string
  toCODECompare: (args: Required<Pick<CODEProps, 'repoPath' | 'diffRefs'>>) => string
  toCODEBranches: (args: Required<Pick<CODEProps, 'repoPath'>>) => string
  toCODEWebhooks: (args: Required<Pick<CODEProps, 'repoPath'>>) => string
  toCODEWebhookNew: (args: Required<Pick<CODEProps, 'repoPath'>>) => string
  toCODEWebhookDetails: (args: Required<Pick<CODEProps, 'repoPath' | 'webhookId'>>) => string
  toCODESettings: (args: Required<Pick<CODEProps, 'repoPath'>>) => string
  toCODECreateWebhook: (args: Required<Pick<CODEProps, 'repoPath'>>) => string
}

export const routes: CODERoutes = {
  toSignIn: (): string => '/signin',
  toSignUp: (): string => '/signup',
  toCODERepositories: ({ space }) => `/${space}`,
  toCODERepository: ({ repoPath, gitRef, resourcePath }) =>
    `/${repoPath}/${gitRef ? '/' + gitRef : ''}${resourcePath ? '/~/' + resourcePath : ''}`,
  toCODEFileEdit: ({ repoPath, gitRef, resourcePath }) => `/${repoPath}/edit/${gitRef}/~/${resourcePath}`,
  toCODECommits: ({ repoPath, commitRef }) => `/${repoPath}/commits/${commitRef}`,
  toCODEPullRequests: ({ repoPath }) => `/${repoPath}/pulls`,
  toCODEPullRequest: ({ repoPath, pullRequestId, pullRequestSection }) =>
    `/${repoPath}/pulls/${pullRequestId}${pullRequestSection ? '/' + pullRequestSection : ''}`,
  toCODECompare: ({ repoPath, diffRefs }) => `/${repoPath}/pulls/compare/${diffRefs}`,
  toCODEBranches: ({ repoPath }) => `/${repoPath}/branches`,
  toCODESettings: ({ repoPath }) => `/${repoPath}/settings`,
  toCODEWebhooks: ({ repoPath }) => `/${repoPath}/webhooks`,
  toCODEWebhookNew: ({ repoPath }) => `/${repoPath}/webhooks/new`,
  toCODEWebhookDetails: ({ repoPath, webhookId }) => `/${repoPath}/webhook/${webhookId}`,
  toCODECreateWebhook: ({ repoPath }) => `/${repoPath}/settings/webhook/new`
}
