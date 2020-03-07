import React from 'react';
import { connect, Dispatch } from 'dva';
import NodeGroup from './components/RuleGroup';
import { RuleGroup } from './typing';
import { Button, Typography, Backdrop, CircularProgress } from '@material-ui/core';
import RuleAddDialog from './components/RuleAddDialog';
import { withSnackbar, WithSnackbarProps } from 'notistack';
import styles from './index.css';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import SyncIcon from '@material-ui/icons/Sync';
import DisplayObject from '@/utils/DisplayObject';

interface NodeGroupsProps extends WithSnackbarProps {
  updatingAll: boolean;
  groups: RuleGroup[];
  dispatch: Dispatch;
}

interface NodeGroupsState {
  ruleAddDialogOpen: boolean;
}

class NodeGroups extends React.Component<NodeGroupsProps, NodeGroupsState> {
  constructor(props: NodeGroupsProps) {
    super(props);
    this.state = {
      ruleAddDialogOpen: false,
    };
  }

  componentWillMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'rules/fetchRules',
    });
  }

  handleRuleAddDialogClose = () => {
    this.setState({ ruleAddDialogOpen: false });
  };

  handleRuleAdd = (group: RuleGroup) => {
    const { dispatch, enqueueSnackbar } = this.props;
    return dispatch({
      type: 'rules/addRule',
      payload: {
        name: group.name,
        proxyGroup: group.proxyGroup,
        url: group.url,
        customRules: group.customRules,
      },
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar(group.name + '添加成功', { variant: 'success' });
          this.handleRuleAddDialogClose();
        } else {
          enqueueSnackbar(msg || group.name + '添加失败', { variant: 'error' });
        }
      },
    });
  };

  handleAllRulesUpdate = () => {
    const { dispatch, enqueueSnackbar } = this.props;
    dispatch({
      type: 'rules/updateAllRules',
      callback: (success: boolean, msg?: string) => {
        if (success) {
          enqueueSnackbar('规则组更新成功', { variant: 'success' });
          this.handleRuleAddDialogClose();
        } else {
          enqueueSnackbar(msg || '规则组更新失败', { variant: 'error' });
        }
      },
    });
  };

  render() {
    const { groups, updatingAll } = this.props;
    const { ruleAddDialogOpen } = this.state;
    return (
      <>
        <RuleAddDialog
          open={ruleAddDialogOpen}
          dialogClose={this.handleRuleAddDialogClose}
          handleRuleAdd={this.handleRuleAdd}
        />
        <div className={styles.globalButtons}>
          <Button
            variant="contained"
            color="secondary"
            startIcon={<SyncIcon />}
            disabled={updatingAll || !groups || !groups.length}
            onClick={this.handleAllRulesUpdate}
          >
            刷新所有规则组
          </Button>
          <Button
            variant="contained"
            color="primary"
            endIcon={<AddCircleOutlineIcon />}
            onClick={() => {
              this.setState({ ruleAddDialogOpen: true });
            }}
          >
            添加规则组
          </Button>
        </div>
        {/* <div><DisplayObject {...this.props.loading} /></div> */}
        {Array.isArray(groups) && groups.length ? (
          groups.map(group => <NodeGroup group={group} updatingAll={updatingAll} />)
        ) : (
          <Typography>还没有规则组，点击上方按钮添加。</Typography>
        )}
      </>
    );
  }
}

export default withSnackbar(
  connect(
    ({
      rules,
      loading,
    }: {
      rules: RuleGroup[];
      loading: {
        effects: { [key: string]: boolean };
      };
    }) => ({
      groups: rules,
      adding: loading.effects['rules/addRule'],
      updatingAll: loading.effects['rules/updateAllRules'],
    }),
  )(NodeGroups),
);
