import React from "react";
import { List, Button } from 'antd';
import reqwest from 'reqwest';
import "./ArticleListPageRollView.less"
import ArticleListItem from "./ArticleListItem";

const count = 3;
const fakeDataUrl = `https://randomuser.me/api/?results=${count}&inc=name,gender,email,nat&noinfo`;

class ArticleListPageRollView extends React.Component {
    constructor(props) {
        super(props);

        this.state =  {
            initLoading: true,
            loading: false,
            data: [],
            list: [],
        };

    }

    componentDidMount() {
        this.getData(res => {
            this.setState({
                initLoading: false,
                data: res.results,
                list: res.results,
            });
        });
    }

    getData = callback => {
        reqwest({
            url: fakeDataUrl,
            type: 'json',
            method: 'get',
            contentType: 'application/json',
            success: res => {
                callback(res);
            },
        });
    };

    onLoadMore = () => {
        this.setState({
            loading: true,
            list: this.state.data.concat([...new Array(count)].map(() => ({ loading: true, name: {} }))),
        });
        this.getData(res => {
            const data = this.state.data.concat(res.results);
            this.setState(
                {
                    data,
                    list: data,
                    loading: false,
                },
                () => {
                    // Resetting window's offsetTop so as to display react-virtualized demo underfloor.
                    // In real scene, you can using public method of react-virtualized:
                    // https://stackoverflow.com/questions/46700726/how-to-use-public-method-updateposition-of-react-virtualized
                    window.dispatchEvent(new Event('resize'));
                },
            );
        });
    };

    render() {
        const { initLoading, loading, list } = this.state;
        const loadMore =
            !initLoading && !loading ? (
                <div
                    style={{
                        textAlign: 'center',
                        marginTop: 12,
                        height: 32,
                        lineHeight: '32px',
                    }}
                >
                    <Button onClick={this.onLoadMore}>loading more</Button>
                </div>
            ) : null;

        return (
            <List
                className="demo-loadmore-list"
                loading={initLoading}
                itemLayout="horizontal"
                loadMore={loadMore}
                dataSource={list}
                renderItem={item => (
                    <ArticleListItem
                        id={item.title}
                        title={item.title}
                        create_time={item.title}
                        pv={item.title}
                        short_content={item.title}
                        picture="https://pic1.zhimg.com/80/v2-af6f3a9444c74d726c63ed5291f9e53d_720w.jpg"
                        description={item.title}/>
                )}
            />
        );
    }

}
export default ArticleListPageRollView;