import List from '@/components/List';
import PageMain from '@/components/PageMain';
import useRequest from '@/hooks/useRequest';
import { postRec } from '@/server/api/posts';
import { DeleteOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Drawer, Modal, Tooltip } from 'antd';
import React, { useState } from 'react';
import { AddPostCate } from './add';
import { PostsChoose } from './posts-choose';

let { confirm } = Modal;
export default function() {
  let { data, load, loading } = useRequest(postRec.list, true);
  let [values, setValues] = useState([{}]);
  let [visible, setVisible] = useState(false);

  let [showChoosePost, setShowChoosePost] = useState(false);
  let [chooseId, setChooseId] = useState(0);

  function delCate(id: number) {
    confirm({
      title: '确定要删除文章？无法撤回',
      icon: <DeleteOutlined color="#d33"></DeleteOutlined>,
      okText: '确定删除',
      okType: 'danger',
      onOk() {
        postRec.del(id).then(() => {
          load();
          setValues([]);
        });
      },
    });
  }

  function editCate(data: any = {}) {
    let { name = '', desc = '', key = '', id = null } = data;
    setValues([
      { name: 'id', value: id },
      { name: 'name', value: name },
      { name: 'desc', value: desc },
      { name: 'key', value: key },
    ]);
    setVisible(true);
  }

  const columns = [
    { title: 'ID', key: 'id', dataIndex: 'id' },
    {
      title: '推荐位名称',
      key: 'name',
      dataIndex: 'name',
    },
    {
      title: '关键词',
      key: 'key',
      dataIndex: 'key',
    },
    { title: '描述', key: 'desc', dataIndex: 'desc' },
    {
      title: '文章统计',
      key: 'count',
      dataIndex: 'count',
      render: (count: number, item: any) => {
        return (
          <div>
            <span>({count}) </span>
            <a
              onClick={() => {
                setShowChoosePost(true);
                setChooseId(item.id);
              }}
            >
              选取
            </a>
          </div>
        );
      },
    },
    {
      title: '更新时间',
      key: 'updated_at',
      dataIndex: 'updated_at',
    },
    {
      title: '创建时间',
      key: 'created_at',
      dataIndex: 'created_at',
    },
  ];

  return (
    <PageMain
      title="推荐位管理"
      subTitle={`共 ${data.length ||
        0} 推荐位，ps：这里需要加一个按规则推荐，最热 最新 活动 或者考虑放在另外的接口`}
    >
      <Drawer width={500} visible={visible} onClose={() => setVisible(false)}>
        <AddPostCate
          onsubmit={() => {
            setVisible(false);
            load();
          }}
          values={values}
        ></AddPostCate>
      </Drawer>

      <Drawer
        width={1200}
        visible={showChoosePost}
        onClose={() => {
          setShowChoosePost(false);
          load();
        }}
      >
        <PostsChoose id={chooseId} />
      </Drawer>

      <List
        onRefresh={load}
        loading={loading}
        leftActions={[
          <Button key="add" type="primary" onClick={editCate}>
            <PlusOutlined />
            <span>添加推荐位</span>
          </Button>,
        ]}
        table={{
          columns: [
            ...columns,
            {
              title: '操作',
              key: 'action',
              dataIndex: 'id',
              render: (id: number, data: any) => {
                return (
                  <div style={{ width: 120 }}>
                    <a onClick={() => editCate(data)}>编辑</a>
                    <span> / </span>
                    {(() => {
                      if (data.count > 0) {
                        return (
                          <Tooltip title="该推荐位下有子文章，不可删除，请先清理文章">
                            <span>不可删除</span>
                          </Tooltip>
                        );
                      } else {
                        return <a onClick={() => delCate(id)}>删除</a>;
                      }
                    })()}
                  </div>
                );
              },
            },
          ],
          dataSource: data instanceof Array ? data : [],
          bordered: true,
          loading,
          rowKey: 'id',
          pagination: {
            total: data.length,
          },
        }}
      />
    </PageMain>
  );
}
